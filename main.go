package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	fc := loadFailConfig()

	// R2: connect + apply migrations BEFORE serving. A failing migration exits
	// non-zero, so a bad migration really breaks the deploy (the modeled risk).
	if err := initDB(context.Background()); err != nil {
		log.Fatalf("inventory-service: db connect failed: %v", err)
	}
	if err := runMigrations(context.Background()); err != nil {
		log.Fatalf("inventory-service: migrations failed: %v", err)
	}

	// "crash" mode: exit shortly after start so Kubernetes CrashLoopBackOffs the
	// pod and ArgoCD health goes Degraded — a real hard-failure outcome.
	if fc.mode == "crash" {
		after := 3
		if s := os.Getenv("CRASH_AFTER_SEC"); s != "" {
			if n, err := strconv.Atoi(s); err == nil {
				after = n
			}
		}
		go func() {
			time.Sleep(time.Duration(after) * time.Second)
			log.Fatalf("FAILURE_MODE=crash: exiting after %ds", after)
		}()
	}

	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())
	mux.HandleFunc("/healthz", instrument("/healthz", failConfig{}, func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
	}))
	// DEEP readiness probe — but inventory-service is a LEAF: it has NO HTTP
	// downstream, so readiness is the DB check ALONE. 200 only if the DB check
	// (SELECT 1) succeeds (or DB is disabled); a broken DB flips this pod
	// NotReady -> ArgoCD Degraded, and because orders-service deep-checks THIS
	// /readyz, the cascade propagates UP the chain (inventory -> orders ->
	// checkout). Liveness stays on shallow /healthz so the kubelet never kills
	// the pod for a DB blip. Left uninstrumented (like /metrics) so frequent
	// probes don't pollute http_requests_total or trip the failure injector.
	mux.HandleFunc("/readyz", func(w http.ResponseWriter, r *http.Request) {
		if !dbHealthy(r.Context()) {
			writeJSON(w, http.StatusServiceUnavailable, map[string]string{"status": "db unavailable"})
			return
		}
		writeJSON(w, http.StatusOK, map[string]string{"status": "ready"})
	})
	registerInventoryRoutes(mux, fc)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("inventory-service listening on :%s (FAILURE_MODE=%q)", port, fc.mode)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatal(err)
	}
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}
