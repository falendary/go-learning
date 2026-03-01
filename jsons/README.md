# About

How to use work with JSON 

# How to run

```go run ./jsons```

```go run ./jsons/error_cases```

# Notes

```
Go struct  --(marshal / encode / serialize)-->  JSON bytes
JSON bytes --(unmarshal / decode / deserialize)-> Go struct
```

The direction is the same.

The vocabulary just comes from different traditions:

"Serialize" → object persistence world

"Encode" → format / protocol world

"Marshal" → systems / RPC world



---

## JSON in REST APIs: Marshal vs Encode/Decode

When building REST APIs in Go, prefer `json.NewEncoder` and `json.NewDecoder` inside HTTP handlers.

For requests:

```go
decoder := json.NewDecoder(r.Body)
decoder.DisallowUnknownFields()
err := decoder.Decode(&req)
```

For responses:

```go
w.Header().Set("Content-Type", "application/json")
json.NewEncoder(w).Encode(resp)
```

This approach streams directly from/to the HTTP layer, avoids unnecessary `[]byte` allocations, and is considered idiomatic for web services.

Use `json.Marshal` and `json.Unmarshal` when working with in-memory data, files, or tests:

```go
b, _ := json.Marshal(user)
json.Unmarshal(b, &user)
```

Rule of thumb:
In HTTP handlers → use Encoder/Decoder.
Outside HTTP → Marshal/Unmarshal.


