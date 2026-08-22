# inventory-service

A small Go HTTP service that reserves inventory (a **leaf** in the rsk0 realistic
stand: it has NO HTTP downstream — it only talks to Postgres). orders-service
deep-checks this service's `/readyz`, so an inventory outage cascades UP the chain
(inventory → orders → checkout).

## Run

```bash
go run .
# http://localhost:8080/healthz
```

## Routes

- `GET  /healthz`   — shallow liveness (process up)
- `GET  /readyz`    — DEEP readiness = Postgres `SELECT 1` only (leaf, no downstream)
- `GET  /metrics`   — Prometheus
- `POST /reserve`   — `{ "sku": "widget", "qty": 3 }` → real INSERT into `inventory`
- `GET  /reserve/{id}`

<!-- ci-verify: trigger build+tag-bump pipeline -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->

<!-- stand: benign copy tweak -->
