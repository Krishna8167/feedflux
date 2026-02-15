 FeedFlux
===========

<p align="center">
  <img 
    src="https://github.com/user-attachments/assets/552f9b36-9f02-4d2e-bad4-267a6cfb6eb9" 
    alt="feedflux Logo" 
    width="380"
  />
</p>

**FeedFlux** is a production-oriented RSS aggregation backend built in Go using Gin.\
The project focuses on clean architecture, background processing, containerized infrastructure, and DevOps-driven development.

This is not just an RSS reader --- it is being developed as a backend systems engineering project.

* * * * *

 Architecture (Current Direction)
-----------------------------------

HTTP (Gin) \
   ↓ \
Handler \
   ↓ \
Service (Business Logic) \
   ↓ \
Repository (Interface) \
   ↓ \
Storage Implementation (In-Memory → PostgreSQL) 


The system follows layered architecture with clear separation of concerns and dependency inversion.

* * * * *

 Project Structure
--------------------

`cmd/
internal/
  handler/
  service/
  repository/
  worker/
  config/
docker-compose.yml
Dockerfile`

* * * * *

 Current Progress
------------------

-   Go module initialized
-   Gin HTTP server with health endpoint
-   Clean layered architecture setup
-   Dockerized PostgreSQL (v16)
-   Docker Compose infrastructure
-   Persistent database volume
-   pgAdmin integration
-   Defined phased backend + DevOps roadmap

* * * * *

 Roadmap
-----------

###  Phase --- Core Backend Architecture

-   Feed CRUD
-   Pagination
-   Validation
-   Structured logging
-   Graceful shutdown
-   In-memory repository implementation

###  Phase --- Database & Data Integrity

-   PostgreSQL repository implementation
-   Proper schema & indexing
-   Unique constraints
-   Idempotent inserts
-   Migrations (golang-migrate)
-   Transaction handling

###  Phase --- Background Worker System

-   RSS polling worker
-   Feed parsing
-   Retry mechanism
-   `last_fetched_at` updates
-   Worker separation (API + Worker services)

###  Phase --- Containerization

-   Multi-stage Dockerfile
-   Non-root execution
-   Docker Compose multi-service setup
-   Healthchecks
-   Restart policies

###  Phase --- Configuration Management

-   Environment-based configuration
-   `.env` support
-   Dev/Prod configs

###  Phase --- Observability

-   Structured logging
-   Prometheus metrics
-   Basic system metrics (feeds, articles, errors, duration)
-   Optional Grafana integration

###  Phase --- Testing

-   Unit tests
-   Repository tests
-   Integration tests
-   CI automation

### Phase --- CI/CD

-   GitHub Actions pipeline
-   Linting
-   Tests
-   Docker image build

###  Phase --- Production Hardening

-   Content timeouts
-   Rate limiting
-   Retry with backoff
-   Security headers
-   Dependency scanning

* * * * *

 Project Goals
----------------

-   Architect a resilient RSS ingestion backend capable of handling multiple feed sources concurrently.
-   Implement background processing patterns
-   Provide reliable data storage with proper constraints, indexing, and transaction handling
-   Expose operational metrics for monitoring ingestion performance and system reliability
-   Build production-aware system design
-   Lay groundwork for future extensibility (authentication, caching, distributed scheduling).

* * * * *

 Running (Development Setup)
------------------------------

Start services:

`docker compose up -d`

Health endpoint:

`GET /health`

* * * * *

 Status
---------

 Active development --- building toward production-grade backend architecture.
