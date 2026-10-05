# Go Todo API

![Go](https://img.shields.io/badge/Go-00ADD8?logo=go&logoColor=white)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-4169E1?logo=postgresql&logoColor=white)
![Docker](https://img.shields.io/badge/Docker-2496ED?logo=docker&logoColor=white)
![Swagger](https://img.shields.io/badge/Swagger-85EA2D?logo=swagger&logoColor=black)

REST API for managing users and tasks, built with **Go** and **PostgreSQL**.

The project uses a feature-based structure with separated **transport, service and repository layers**, database migrations, middleware, Swagger documentation and Docker-based infrastructure.

## Features

- Users CRUD
- Tasks CRUD
- Pagination and task filtering
- Task statistics with optional user and date filters
- Partial updates with `PATCH`
- Request validation
- PostgreSQL persistence with `pgx`
- Database migrations with `golang-migrate`
- Optimistic locking using entity versions
- API versioning under `/api/v1`
- Swagger / OpenAPI documentation
- CORS, request ID, logging and panic recovery middleware
- Graceful HTTP server shutdown
- Structured logging with Zap
- Docker and Docker Compose
- Simple web interface

## Tech Stack

**Backend**

- Go
- `net/http`
- `pgx`
- `go-playground/validator`
- Zap
- envconfig

**Infrastructure**

- PostgreSQL
- Docker
- Docker Compose
- golang-migrate

**Documentation**

- swag
- http-swagger

## Architecture

Each main feature is split into three layers:

```text
HTTP Request
     │
     ▼
 Transport
     │
     ▼
  Service
     │
     ▼
Repository
     │
     ▼
PostgreSQL
```

- **Transport** — HTTP handlers, request parsing, validation and responses
- **Service** — application and business logic
- **Repository** — persistence and SQL queries

Dependencies between layers are defined through interfaces.

## Project Structure

```text
.
├── cmd/
│   └── todoapp/
│       ├── main.go
│       └── Dockerfile
│
├── internal/
│   ├── core/
│   │   ├── config/
│   │   ├── domain/
│   │   ├── errors/
│   │   ├── logger/
│   │   ├── repository/
│   │   └── transport/
│   │
│   └── features/
│       ├── users/
│       ├── tasks/
│       ├── statistics/
│       └── web/
│
├── migrations/
├── docs/
├── public/
├── docker-compose.yml
├── Makefile
└── .env.example
```

Each feature contains its own `transport`, `service` and `repository` packages where needed.

## API

Base URL:

```text
http://127.0.0.1:5050/api/v1
```

### Users

| Method | Endpoint | Description |
|---|---|---|
| `POST` | `/users` | Create user |
| `GET` | `/users` | Get users |
| `GET` | `/users/{id}` | Get user by ID |
| `PATCH` | `/users/{id}` | Update user |
| `DELETE` | `/users/{id}` | Delete user |

Pagination:

```text
GET /api/v1/users?limit=20&offset=0
```

### Tasks

| Method | Endpoint | Description |
|---|---|---|
| `POST` | `/tasks` | Create task |
| `GET` | `/tasks` | Get tasks |
| `GET` | `/tasks/{id}` | Get task by ID |
| `PATCH` | `/tasks/{id}` | Update task |
| `DELETE` | `/tasks/{id}` | Delete task |

Filtering and pagination:

```text
GET /api/v1/tasks?user_id=5&limit=20&offset=0
```

### Statistics

```text
GET /api/v1/statistics
```

Optional filters:

```text
GET /api/v1/statistics?user_id=5&from=2026-10-01&to=2026-11-01
```

## Swagger

Interactive API documentation is available at:

```text
http://127.0.0.1:5050/swagger/
```

Regenerate Swagger files:

```bash
make swagger-gen
```

## Getting Started

### Requirements

- Go
- Docker
- Docker Compose
- Make

### Clone the repository

```bash
git clone https://github.com/jabrail059/golang-todoapp.git
cd golang-todoapp
```

### Configure environment

Create `.env` from the example:

```bash
cp .env.example .env
```

Then configure the required values:

```env
HTTP_ADDR=:5050
HTTP_SHUTDOWN_TIMEOUT=30s
ALLOWED_ORIGINS=http://localhost:5050

POSTGRES_USER=todoapp
POSTGRES_PASSWORD=your_password
POSTGRES_DB=todoapp
POSTGRES_TIMEOUT=10s

LOGGER_LEVEL=DEBUG
TIME_ZONE=UTC
```

## Run Locally

Start PostgreSQL:

```bash
make env-up
```

Expose PostgreSQL on `localhost:5432`:

```bash
make env-port-forward
```

Apply migrations:

```bash
make migrate-up
```

Run the application:

```bash
make todoapp-run
```

The server will be available at:

```text
http://127.0.0.1:5050
```

## Run with Docker

Build and start the application:

```bash
make todoapp-deploy
```

Apply migrations:

```bash
make migrate-up
```

Check running services:

```bash
make ps
```

Stop the application:

```bash
make todoapp-undeploy
```

## Useful Commands

| Command | Description |
|---|---|
| `make env-up` | Start PostgreSQL |
| `make env-down` | Stop PostgreSQL |
| `make migrate-up` | Apply migrations |
| `make migrate-down` | Roll back migration |
| `make migrate-create seq=name` | Create a new migration |
| `make todoapp-run` | Run application locally |
| `make todoapp-deploy` | Build and run with Docker |
| `make todoapp-undeploy` | Stop Docker deployment |
| `make swagger-gen` | Regenerate Swagger docs |
| `make ps` | Show Compose services |

## Web Interface

A simple web interface is also served by the application:

```text
http://127.0.0.1:5050/
```