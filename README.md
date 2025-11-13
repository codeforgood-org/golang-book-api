# Book API

A production-ready RESTful API for managing books, built with Go and following best practices for project structure and organization.

## Features

- **RESTful API** with full CRUD operations for books
- **Clean Architecture** with organized package structure
- **Middleware Support** including logging, CORS, and panic recovery
- **Comprehensive Testing** with unit tests
- **Configuration Management** via environment variables
- **Health Check Endpoint** for monitoring
- **Docker Support** with multi-stage builds
- **CI/CD Pipeline** with GitHub Actions
- **Thread-Safe** in-memory storage with mutex protection

## Project Structure

```
.
├── cmd/
│   └── api/
│       └── main.go              # Application entry point
├── internal/
│   ├── config/
│   │   └── config.go            # Configuration management
│   ├── handlers/
│   │   ├── books.go             # Book HTTP handlers
│   │   └── health.go            # Health check handler
│   ├── middleware/
│   │   ├── cors.go              # CORS middleware
│   │   ├── logger.go            # Request logging middleware
│   │   └── recovery.go          # Panic recovery middleware
│   ├── models/
│   │   ├── book.go              # Book model and validation
│   │   ├── book_test.go         # Book model tests
│   │   └── errors.go            # Domain errors
│   └── storage/
│       ├── storage.go           # Storage interface
│       ├── memory.go            # In-memory implementation
│       └── memory_test.go       # Storage tests
├── pkg/
│   └── logger/
│       └── logger.go            # Logging utilities
├── .github/
│   └── workflows/
│       └── ci.yml               # CI/CD pipeline
├── Dockerfile                   # Multi-stage Docker build
├── docker-compose.yml           # Docker Compose configuration
├── Makefile                     # Common development tasks
├── .env.example                 # Example environment variables
├── go.mod                       # Go module definition
└── README.md                    # This file
```

## API Endpoints

### Health Check
- `GET /health` - Check API health status

### Books
- `GET /books` - Get all books
- `POST /books` - Create a new book
- `GET /books/{id}` - Get a book by ID
- `DELETE /books/{id}` - Delete a book by ID

## Getting Started

### Prerequisites

- Go 1.21 or higher
- Docker (optional)
- Make (optional)

### Installation

1. Clone the repository:
```bash
git clone https://github.com/codeforgood-org/golang-book-api.git
cd golang-book-api
```

2. Install dependencies:
```bash
go mod download
```

3. Copy the example environment file:
```bash
cp .env.example .env
```

4. Run the application:
```bash
go run cmd/api/main.go
```

The server will start on `http://localhost:8080`

### Using Make

The project includes a Makefile with common tasks:

```bash
make build      # Build the application
make run        # Run the application
make test       # Run tests
make test-cover # Run tests with coverage
make lint       # Run linter
make clean      # Clean build artifacts
make docker     # Build Docker image
make help       # Show available commands
```

### Using Docker

Build and run with Docker:

```bash
docker build -t book-api .
docker run -p 8080:8080 book-api
```

Or use Docker Compose:

```bash
docker-compose up
```

## Usage Examples

### Create a Book

```bash
curl -X POST http://localhost:8080/books \
  -H "Content-Type: application/json" \
  -d '{
    "title": "The Go Programming Language",
    "author": "Alan A. A. Donovan"
  }'
```

Response:
```json
{
  "id": 123456,
  "title": "The Go Programming Language",
  "author": "Alan A. A. Donovan"
}
```

### Get All Books

```bash
curl http://localhost:8080/books
```

Response:
```json
[
  {
    "id": 123456,
    "title": "The Go Programming Language",
    "author": "Alan A. A. Donovan"
  }
]
```

### Get a Book by ID

```bash
curl http://localhost:8080/books/123456
```

### Delete a Book

```bash
curl -X DELETE http://localhost:8080/books/123456
```

### Health Check

```bash
curl http://localhost:8080/health
```

Response:
```json
{
  "status": "ok"
}
```

## Configuration

The application can be configured using environment variables:

| Variable | Description | Default |
|----------|-------------|---------|
| `SERVER_PORT` | Port to run the server on | `8080` |
| `LOG_LEVEL` | Logging level (info, warning, error) | `info` |

## Testing

Run all tests:
```bash
go test ./...
```

Run tests with coverage:
```bash
go test -cover ./...
```

Run tests with detailed coverage report:
```bash
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

## Development

### Code Organization

- **`cmd/`** - Application entry points
- **`internal/`** - Private application code
  - **`config/`** - Configuration management
  - **`handlers/`** - HTTP request handlers
  - **`middleware/`** - HTTP middleware
  - **`models/`** - Domain models and business logic
  - **`storage/`** - Data storage layer
- **`pkg/`** - Public packages that can be imported by other projects

### Adding a New Endpoint

1. Define the model in `internal/models/`
2. Add storage methods in `internal/storage/`
3. Create handler in `internal/handlers/`
4. Register route in `cmd/api/main.go`
5. Add tests

### Best Practices

- Keep packages focused and cohesive
- Use interfaces for dependencies
- Write tests for all business logic
- Use meaningful error messages
- Log important events
- Handle errors appropriately
- Use middleware for cross-cutting concerns

## CI/CD

The project uses GitHub Actions for continuous integration. On every push:

1. Code is checked out
2. Dependencies are installed
3. Tests are run
4. Code is built
5. Docker image is built (optional)

See `.github/workflows/ci.yml` for details.

## Future Enhancements

- [ ] Database integration (PostgreSQL, MongoDB)
- [ ] Authentication and authorization
- [ ] API documentation with Swagger/OpenAPI
- [ ] Rate limiting
- [ ] Caching layer
- [ ] Pagination for GET /books
- [ ] Search and filtering
- [ ] Metrics and monitoring (Prometheus)
- [ ] GraphQL support

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contact

Project Link: [https://github.com/codeforgood-org/golang-book-api](https://github.com/codeforgood-org/golang-book-api)
