ecommerce-order-service/
├── cmd/                      # Entrypoints for different services
│   ├── order-service/        # main.go for order service
│   ├── inventory-service/    # main.go for inventory service
│   └── notification-service/ # main.go for notification service
│
├── internal/                 # Private app code (domain + business logic)
│   ├── orders/               # Order domain logic
│   │   ├── handler.go        # HTTP handlers (net/http)
│   │   ├── service.go        # Business logic (use cases)
│   │   └── repository.go     # Data access (calls sqlc)
│   │
│   ├── inventory/            # Inventory domain logic
│   │   ├── handler.go
│   │   ├── service.go
│   │   └── repository.go
│   │
│   └── notification/         # Notification domain logic
│       ├── handler.go
│       ├── service.go
│       └── repository.go
│
├── pkg/                      # Shared utilities (JWT, middlewares, etc.)
│   ├── auth/
│   ├── config/
│   ├── kafka/
│   └── redis/
│
├── db/
│   ├── migration/            # SQL migration files (up/down)
│   └── sqlc/                 # Auto-generated sqlc code
│       ├── db.go             # DB connection helpers
│       ├── models.go         # Types
│       ├── inventory.sql.go
│       └── orders.sql.go
│
├── api/
│   ├── swagger/              # OpenAPI/Swagger specs
│   └── docs/                 # Generated Swagger docs
│
├── test/                     # Integration / E2E tests
│   ├── orders_test.go
│   └── inventory_test.go
│
├── docker/                   # Docker-specific config (compose files, init scripts)
│   └── docker-compose.yml
│
├── Makefile                  # Shortcuts (run tests, migrate, etc.)
├── go.mod
└── go.sum
