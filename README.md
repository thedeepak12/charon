# Charon

A lightweight, thread-safe token bucket rate limiter implementation in Go with HTTP middleware support.

## Features

- **Token Bucket Algorithm**: Core rate limiting logic with configurable capacity and refill rate.
- **Thread-Safe**: Uses mutex locks for concurrent access.
- **Flexible Interface**: Limiter interface for easy extensibility.
- **In-Memory Storage**: Manage multiple rate limiters by key (e.g., IP address).
- **HTTP Middleware**: Ready-to-use middleware for Go HTTP servers.
- **Per-IP Rate Limiting**: Automatic IP extraction with support for `X-Forwarded-For` and `X-Real-IP` headers.
- **Rate Limit Headers**: Includes `X-RateLimit-Limit`, `X-RateLimit-Remaining`, and `Retry-After` headers.
- **Request Logging**: Built-in logging for monitoring rate limit decisions.

## Tech Stack

- **Language**: Go 1.26.4
- **Libraries**: net/http, sync, time

## Project Structure

```text
charon/
├── cmd/
│   └── demo/
│       └── main.go              # Demo HTTP server with rate limiting
├── internal/
│   ├── middleware/
│   │   └── ratelimit.go         # HTTP middleware for rate limiting
│   └── ratelimiter/
│       ├── config.go            # Configuration and validation
│       ├── interface.go         # Limiter interface definition
│       ├── storage.go           # In-memory storage for multiple limiters
│       └── token_bucket.go      # Token bucket algorithm implementation
└── go.mod                       # Go module descriptor
```

## Setup

1. Clone the repository:
```bash
git clone https://github.com/thedeepak12/charon.git
cd charon
```

2. Tidy the Go modules:
```bash
go mod tidy
```

3. Run the demo application:
```bash
go run ./cmd/demo/main.go
```

The demo server will start on port 4000 and demonstrate rate limiting per IP address.

4. Test the rate limiter:

**Using curl:**
```bash
# Send a single request to verify that the server is responding
curl -i http://localhost:4000

# Send multiple requests quickly to trigger rate limiting
for i in {1..100}; do curl -i http://localhost:4000/; done
```

**Or manually in your browser:**
- Open `http://localhost:4000` in your browser.
- Refresh the page multiple times to observe rate limiting in action.

**Responses:**
- **200**: Request allowed
- **429**: Rate limit exceeded

## License

Distributed under the MIT License. See [LICENSE](https://github.com/thedeepak12/charon/blob/main/LICENSE) for more information.
