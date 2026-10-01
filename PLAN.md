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
