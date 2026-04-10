# Demo Payment Service

Payment processing microservice for the Gateway e-commerce demo. Built with Go and Gin.

## Quick Start

```bash
go run main.go   # runs on port 8082
```

## API Endpoints

| Method | Path | Description |
|--------|------|-------------|
| GET | /health | Health check |
| GET | /metrics | Metrics (Gateway-compatible) |
| POST | /api/v1/payments | Process payment (5% mock failure rate) |
| GET | /api/v1/payments/:id | Get payment by ID |
| GET | /api/v1/payments | List payments |
| POST | /api/v1/payments/:id/refund | Refund payment |

## Payment Processing

Mock payment processor with configurable 5% failure rate and 200-500ms simulated latency.
