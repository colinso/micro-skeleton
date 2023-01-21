# Micro-Skeleton

This repo is a skeleton project that can be forked to spin up a simple API microservice in AWS. It has several example endpoints to use to create a backend for a storefront.

Dependencies:
* Apex (Structured Logging) - https://github.com/apex/log
* Wire (Dependency Injection) - https://github.com/google/wire
* Chi (HTTP Router) - https://github.com/go-chi/chi
* Goose (DB Migrations) - https://github.com/pressly/goose
* PostgreSQL/PGX (DB Driver) - https://github.com/jackc/pgx

**TODOS**
- Swagger documentation (https://github.com/swaggo/swag)
- Improved logging and other monitoring
- Endpoint validation (https://tutorialedge.net/golang/validating-http-json-requests/)
- Add testing and mock generation
- Context propogation
