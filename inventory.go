package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"
)

// Reservation is one inventory reservation: a quantity of a SKU held for an order.
type Reservation struct {
	ID  string `json:"id"`
	SKU string `json:"sku"`
	Qty int    `json:"qty"`
}

var (
	mu           sync.Mutex
	reservations = map[string]*Reservation{}
	nextID       = 1
)

// Available returns the units still available given a stock level and a reserved
// quantity. Never negative.
func Available(stock, reserved int) int {
	if reserved >= stock {
		return 0
	}
	return stock - reserved
}

func registerInventoryRoutes(mux *http.ServeMux, fc failConfig) {
	mux.HandleFunc("/reserve", instrument("/reserve", fc, businessGate(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		var in struct {
			SKU string `json:"sku"`
			Qty int    `json:"qty"`
		}
		if err := json.NewDecoder(r.Body).Decode(&in); err != nil || in.Qty <= 0 {
			http.Error(w, "invalid reservation", http.StatusBadRequest)
			return
		}
		// R2: with a DB, do a REAL INSERT+SELECT against the inventory table; a DB
		// error surfaces as a real 500 (so a broken schema is a real failure).
		if dbEnabled() {
			res, err := dbCreateReservation(r.Context(), in.SKU, in.Qty)
			if err != nil {
				http.Error(w, "reservation persistence failed", http.StatusInternalServerError)
				return
			}
			writeJSON(w, http.StatusCreated, res)
			return
		}
		mu.Lock()
		id := fmt.Sprintf("inv_%d", nextID)
		nextID++
		res := &Reservation{ID: id, SKU: in.SKU, Qty: in.Qty}
		reservations[id] = res
		mu.Unlock()
		writeJSON(w, http.StatusCreated, res)
	})))

	mux.HandleFunc("/reserve/", instrument("/reserve/", fc, businessGate(func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimPrefix(r.URL.Path, "/reserve/")
		if dbEnabled() {
			res, err := dbGetReservation(r.Context(), id)
			if err != nil {
				http.Error(w, "reservation lookup failed", http.StatusInternalServerError)
				return
			}
			if res == nil {
				http.Error(w, "not found", http.StatusNotFound)
				return
			}
			writeJSON(w, http.StatusOK, res)
			return
		}
		mu.Lock()
		res, ok := reservations[id]
		mu.Unlock()
		if !ok {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}
		writeJSON(w, http.StatusOK, res)
	})))
}
