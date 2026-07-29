package main

// inventory-service is a LEAF in the dependency mesh: it has NO HTTP downstream
// (it only talks to Postgres). So there is no downstreamGate / cascade check
// here — unlike orders-service, which deep-checks payments + inventory. The only
// per-request work the business gate does is R3 footprint accumulation.

import "net/http"

// businessGate wraps a business handler and, on the business path ONLY (never
// /healthz, /readyz or /metrics), accumulates the real bounded memory footprint.
// Under sustained load RSS plateaus at ~MEM_FOOTPRINT_MB — so a risky change that
// lowers resources.limits.memory below that gets the container OOMKilled.
func businessGate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		recordFootprint()
		next(w, r)
	}
}
