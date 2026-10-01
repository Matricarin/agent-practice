# agent-practice

Minimal Go HTTP backend that adds two numbers.

## Run

```bash
go run .
```

Server listens on `:8080`.

## API

### `POST /sum`

Request:

```json
{ "a": 2, "b": 3 }
```

Response:

```json
{ "sum": 5 }
```

### Example

```bash
curl -s -X POST http://localhost:8080/sum \
  -H 'Content-Type: application/json' \
  -d '{"a":2,"b":3}'
```
