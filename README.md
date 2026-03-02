# Go Learning

Structured practice repository for learning Go step by step

The goal is to understand the language, concurrency model, standard library, and how to build production-ready applications

---

## Repository Structure

```
go-learning/
│
├── basics/
├── errors/            - how to work with errors and how to define them
├── files/             - how to work with files (create, delete, read line by line)
├── generics/
├── http/              - how to start basic http server and serve routes
├── http-client/       - how to make http requests (server-to-server)
├── interfaces/
├── io/                - io reader/writer examples
├── jsons/             - how to work with JSON (marshal/unmarshal, encode/decode)
├── slices/            - how to use Arrays, Lists, Slices
├── structs/           - how to use Dicts / Classes
├── tests/             - how to write tests
└── README.md
```

Each folder contains

* Source code
* Tests
* Short explanation of the topic

---


## Install Go

Download Go from [https://go.dev/dl/](https://go.dev/dl/) and install it using the installer for your operating system. After installation, verify it works by running:

### macOS
```brew install go```

### Linux
```
wget https://go.dev/dl/go1.22.0.linux-amd64.tar.gz
sudo rm -rf /usr/local/go
sudo tar -C /usr/local -xzf go1.22.0.linux-amd64.tar.gz
```
---

Verify installation

```bash
go version
```

You should see the installed version printed in the terminal.

## How to Run Code

Run an example

```bash
go run ./basics
```

Run all tests

```bash
go test ./...
```

Run with race detector

```bash
go test -race ./...
```

Format code

```bash
go fmt ./...
```

Clean dependencies

```bash
go mod tidy
```

---

## Phase 1 — Language Core

Focus on fundamentals

* Pointers and value semantics
* Structs and JSON tags
* Methods and receivers
* Interfaces
* Error handling
* Generics

---

## Phase 2 — Concurrency

Focus on orchestration

* Goroutines
* Channels
* Select statement
* Worker pool
* Fan-in and fan-out
* Context and cancellation
* Sync primitives

---

## Phase 3 — Standard Library

Focus on core tools

* IO
* HTTP
* JSON encoding
* Testing
* Benchmarks

---

## Phase 4 — Real Application

Build a structured REST API

* Project layout
* Dependency injection
* Middleware
* Logging
* Database integration
* Graceful shutdown
* Profiling

---

## Phase 5 — Applied Concurrency

Build a concurrent data pipeline

* Worker pool
* Backpressure
* Cancellation
* Error propagation
* Clean shutdown

---

