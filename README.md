# ⚒️ Simple Go API

This is a simple Go API I'm learning with. This isn't something I'd currently use for a production system.

## 📋 Requirements

- [Go](https://go.dev/) - Version 1.24.2+

## 🚀 Running the Project

```bash
go run *.go
```

The project will be running at [http://localhost:8080](http://localhost:8080).

### Available Routes

- POST [http://localhost:8080/item](http://localhost:8080/item)

__Example Payload__

```bash
curl -X POST -H "Content-Type: application/json" -d '{"name": "New Item", "value": 30}' http://localhost:8080/item
```

