# Go Hello World Project Plan

## Overview

Build a minimal Go CLI that prints `Hello, World!` using only the standard library.

## Layout

Place the module at the repo root (matches the existing Go-oriented `.gitignore`):

```
.
├── go.mod
├── main.go
├── PLAN.md
├── README.md
├── LICENSE
└── .gitignore
```

## Steps

1. **Initialize the module**

   ```bash
   go mod init agent-practice
   ```

2. **Add `main.go`**

   ```go
   package main

   import "fmt"

   func main() {
       fmt.Println("Hello, World!")
   }
   ```

3. **Run**

   ```bash
   go run .
   ```

4. **Optional: build a binary**

   ```bash
   go build -o hello
   ```

## Verify

- `go run .` prints `Hello, World!` to stdout.
- The built binary (`hello`) is ignored by `.gitignore` (covered by binary patterns / not committed).

## Out of scope

- Tests, CI, and third-party dependencies

---

# Upgrade: Simple Sum Backend

## Overview

Upgrade the project from a CLI hello-world app to a minimal HTTP backend. Accept a POST with two numbers and return a JSON response with their sum. Use only the Go standard library (`net/http`, `encoding/json`).

## Layout

Keep the module at the repo root:

```
.
├── go.mod
├── main.go
├── PLAN.md
├── README.md
├── LICENSE
└── .gitignore
```

## API

- **Method / path:** `POST /sum`
- **Request body (JSON):**

  ```json
  { "a": 2, "b": 3 }
  ```

- **Success response (JSON):**

  ```json
  { "sum": 5 }
  ```

- **Errors:** return a non-2xx status with a short JSON error body when the method is wrong, the body is invalid, or fields are missing/not numbers.

## Steps

1. **Replace `main.go` with an HTTP server**

   - Listen on `:8080` (or another fixed port).
   - Register `POST /sum`.
   - Decode JSON `{ "a": number, "b": number }`.
   - Respond with JSON `{ "sum": a + b }` and `Content-Type: application/json`.

2. **Run**

   ```bash
   go run .
   ```

3. **Smoke-test with curl**

   ```bash
   curl -s -X POST http://localhost:8080/sum \
     -H 'Content-Type: application/json' \
     -d '{"a":2,"b":3}'
   ```

   Expected output: `{"sum":5}`

## Verify

- Server starts without errors.
- Valid POST returns JSON with the correct sum.
- Non-POST methods and bad JSON are rejected with an error response.

## Out of scope

- Auth, persistence, frameworks, tests, and CI
