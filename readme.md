crypto_exchange/
├── cmd/
│   └── app/
│       └── main.go
├── internal/
│   ├── api/
│   │   ├── handlers/
│   │   │   ├── auth.go
│   │   │   ├── trades.go
│   │   │   ├── quotes.go
│   │   │   └── transactions.go
│   │   ├── middleware/
│   │   │   ├── auth_middleware.go
│   │   │   └── logging.ноgo
│   │   └── router.go
│   ├── config/
│   │   └── config.go
│   ├── database/
│   │   ├── migrations/
│   │   │   └── *.sql
│   │   └── database.go
│   ├── models/
│   │   ├── user.go
│   │   ├── order.go
│   │   ├── transaction.go
│   │   └── quote.go
│   ├── repository/
│   │   ├── user_repository.go
│   │   ├── order_repository.go
│   │   ├── transaction_repository.go
│   │   └── quote_repository.go
│   ├── service/
│   │   ├── auth_service.go
│   │   ├── trade_service.go
│   │   ├── quote_service.go
│   │   └── transaction_service.go
│   └── utils/
│       ├── jwt.go
│       └── hash.go
├── pkg/
│   └── external_api/
│       └── crypto_api.go
├── migrations/
│   └── *.sql
├── Dockerfile
├── docker-compose.yml
├── .env
├── go.mod
├── go.sum
└── README.md
