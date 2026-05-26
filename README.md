# E-Commerce API

![Go](https://img.shields.io/badge/Go-1.21-00ADD8?style=flat-square&logo=go&logoColor=white)
![Gin](https://img.shields.io/badge/Gin-1.10-00ADD8?style=flat-square)
![MySQL](https://img.shields.io/badge/MySQL-8.0-4479A1?style=flat-square&logo=mysql&logoColor=white)
![Docker](https://img.shields.io/badge/Docker-Ready-2496ED?style=flat-square&logo=docker&logoColor=white)
![Swagger](https://img.shields.io/badge/Swagger-Docs-85EA2D?style=flat-square&logo=swagger&logoColor=black)

E-commerce REST API built with Go, Clean Architecture, and [go-toolkit](https://github.com/alpardfm/go-toolkit) — a personal utility library used across multiple projects.

> 🚧 **Status:** In active development. Dashboard APIs complete, mobile APIs in progress.

---

## Highlights

- **Clean Architecture** — handler → usecase → domain → database, with interfaces at every boundary
- **Powered by [go-toolkit](https://github.com/alpardfm/go-toolkit)** — uses own library for SQL, logging, JWT, config, error handling, query building
- **Role-based access** — admin, librarian, member with JWT authorization
- **Location-based auth** — dashboard login validates coordinates via Haversine distance
- **Swagger docs** — auto-generated API documentation
- **12 domain modules** — users, products, categories, cart, orders, payments, refund, reviews, OTP, role, location, order items

---

## Architecture

```
┌──────────────────────────────────────────────────────┐
│                   Handler (REST)                       │
│  Gin Router → Middleware → Handler                    │
├──────────────────────────────────────────────────────┤
│                  Usecase Layer                         │
│  Business logic, JWT validation, authorization        │
├──────────────────────────────────────────────────────┤
│                  Domain Layer                          │
│  Repository pattern, SQL queries, transactions        │
├──────────────────────────────────────────────────────┤
│                    MySQL 8.0                           │
└──────────────────────────────────────────────────────┘
```

```
e-commerce/
├── src/
│   ├── cmd/main.go              # Application entrypoint
│   ├── business/
│   │   ├── domain/              # Data access layer (12 domains)
│   │   └── usecase/             # Business logic layer
│   ├── entity/                  # Data models & DTOs
│   ├── handler/rest/            # HTTP handlers + routes
│   ├── middlewares/             # Request context middleware
│   └── utils/                   # Config, helpers, keys
├── docs/
│   ├── sql/schema.sql           # Full database schema
│   └── swagger/                 # Auto-generated Swagger
├── etc/cfg/                     # Runtime configuration
├── Dockerfile                   # Multi-stage build
└── docker-compose.yaml          # App + MySQL
```

---

## Tech Stack

| Layer | Technology |
|-------|-----------|
| Language | Go 1.21 |
| Framework | Gin |
| Database | MySQL 8.0 |
| ORM/Query | go-toolkit/sql + go-toolkit/query |
| Auth | JWT (go-toolkit/tokens) |
| Logging | Zerolog (go-toolkit/log) |
| Config | Viper (go-toolkit/configreader) |
| Docs | Swagger (swaggo) |
| Container | Docker + docker-compose |

---

## go-toolkit Integration

This project uses [go-toolkit](https://github.com/alpardfm/go-toolkit) extensively:

| Package | Usage |
|---------|-------|
| `sql` | Database connection (leader/follower pattern) |
| `log` | Structured logging |
| `tokens` | JWT creation & validation (generics-based) |
| `errors` + `codes` | Error wrapping with HTTP status mapping |
| `query` | Dynamic SQL query builder from struct tags |
| `configreader` | Viper-based config loading |
| `configbuilder` | AWS SSM parameter store config generation |
| `distance` | Haversine distance calculation for location auth |
| `appcontext` | Request-scoped context values |
| `parser` | JSON marshaling |

---

## API Endpoints

### Implemented ✅

| Method | Path | Description |
|--------|------|-------------|
| `POST` | `/api/loginDashboard` | Dashboard login (email + password + location) |
| `GET` | `/api/pagination/categories` | List categories (paginated) |
| `GET/POST/PUT/DELETE` | `/api/categories/:id` | CRUD categories |
| `GET` | `/api/pagination/location` | List locations (paginated) |
| `GET/POST/PUT/DELETE` | `/api/location/:id` | CRUD locations |
| `GET` | `/api/pagination/role` | List roles (paginated) |
| `GET/POST/PUT/DELETE` | `/api/role/:id` | CRUD roles |
| `GET` | `/swagger/*` | Swagger UI |
| `GET` | `/ping` | Health check |

### Planned 🚧

- Products CRUD
- Cart CRUD
- Orders + Payments flow
- Mobile auth (register, OTP, pincode)
- Reviews, Refund

---

## Database Schema

11 tables with full audit trail (`created_at`, `created_by`, `updated_at`, `updated_by`, soft delete):

`users` · `otp` · `role` · `location` · `categories` · `products` · `orders` · `order_items` · `cart` · `reviews` · `payments` · `refund`

Full schema: [`docs/sql/schema.sql`](docs/sql/schema.sql)

---

## Quick Start

### Prerequisites
- Go 1.21+
- Docker & Docker Compose

### 1. Clone & Setup
```bash
git clone https://github.com/alpardfm/e-commerce.git
cd e-commerce
```

### 2. Start Database
```bash
docker compose up -d db
```

### 3. Run API
```bash
go run ./src/cmd/main.go
```

API available at `http://localhost:3001`. Swagger at `/swagger/`.

---

## License

MIT
