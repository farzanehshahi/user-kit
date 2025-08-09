# user-kit

A production-ready **User CRUD** microservice written in **Go** with **go-kit**, fully instrumented for **Prometheus** and **Grafana**.

- Developed based on go-kit framework
- Structured logging & middlewares
- One-command local stack via Docker Compose

## Table of Contents

- [Architecture](#architecture)
- [Directory Layout](#directory-layout)
- [Quick Start (Docker)](#quick-start-docker)
- [Configuration](#configuration)
- [Database Migrations](#database-migrations)
- [API Endpoints](#api-endpoints)
- [Monitoring](#monitoring)
- [Local Development (without Docker)](#local-development-without-docker)
- [Testing](#testing)
- [Roadmap](#roadmap)
- [License](#license)

## Architecture

This service follows go-kit’s classic layering:

- **transport/** (HTTP): route handlers + request/response encoding/decoding  
- **endpoint/**: endpoint definitions and cross-cutting middlewares (metrics, logging, timeout)  
- **service/**: domain logic (create/read/update/delete users)  
- **repository/**: database access (PostgreSQL)

This separation keeps business logic testable and swaps transports or stores without touching the core.

## Directory Layout

```text
.
├─ cmd/                 # main application wiring
├─ config/              # application configs (YAML/env)
├─ internal/            # service, endpoint, transport, repository
├─ migration/           # SQL migrations
├─ pkg/                 # shared helpers (logger, config, errors, ...)
├─ prometheus/          # prometheus.yml and scrape config
├─ Dockerfile
├─ docker-compose.yml
├─ entrypoint.sh
└─ wait-for
```

## Quick Start (Docker)

**Prereqs:** Docker & Docker Compose.

```bash
# Build and start the full stack 
docker compose up -d --build

# Follow app logs
docker compose logs -f app
```

**Default ports (adjust if your compose uses different ones):**

- API: `http://localhost:8080`
- PostgreSQL: `localhost:5432`
- Prometheus: `http://localhost:9090`
- Grafana: `http://localhost:3000` (default creds: `admin` / `admin` unless overridden)

## API Endpoints

> Endpoints below follow common CRUD patterns. If your actual routes differ, update accordingly.

- `POST   /users` — Create a user  
  Request:
  ```json
  { "name": "Alice", "email": "alice@example.com" }
  ```
- `GET    /users` — List users (supports `?page=&limit=` if implemented)
- `GET    /users/{id}` — Get user by ID
- `PUT    /users/{id}` — Update user
- `DELETE /users/{id}` — Delete user
- `GET    /health` — Liveness/health check
- `GET    /metrics` — Prometheus metrics

## Monitoring

- **Metrics**: exported at `/metrics` (standard Go & go-kit counters/histograms like request count, latency, in-flight).
- **Prometheus**: point a job to the app’s `/metrics` target.  
- **Grafana**: included in Docker Compose. Add Prometheus as a data source:
  - Inside Docker network: `http://prometheus:9090`
  - From host: `http://localhost:9090`

## Roadmap

- [ ] OpenAPI/Swagger docs
- [ ] Auth (JWT/OIDC)

## License

MIT (or whatever your `LICENSE` file states).
