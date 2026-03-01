
# About

How to use http server

- Example how to create route /health
- Response with sample JSON data

# How to run

```go run ./http```

# How to test

Echo endpoint
```
curl -X POST http://localhost:8080/echo \
  -H "Content-Type: application/json" \
  -d '{"message":"hello"}'
```

Health endpoint
```
curl -X GET http://localhost:8080/health \
  -H "Content-Type: application/json" 
```

# How to run tests

```go test ./http -v ```