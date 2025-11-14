# Book API

A production-ready RESTful API for managing books, built with Go and following best practices for project structure and organization.

## Features

- **RESTful API** with full CRUD operations (Create, Read, Update, Delete)
- **Pagination** with configurable page size (up to 100 items per page)
- **Filtering & Search** by title, author, or both
- **Request ID Tracking** for distributed tracing
- **Clean Architecture** with organized package structure
- **Middleware Support** including request ID, logging, CORS, and panic recovery
- **Comprehensive Testing** with unit tests and benchmarks
- **Configuration Management** via environment variables
- **OpenAPI/Swagger Specification** for API documentation
- **Health Check Endpoint** for monitoring
- **Docker Support** with multi-stage builds
- **CI/CD Pipeline** with GitHub Actions
- **Thread-Safe** in-memory storage with mutex protection
- **Sample Data Seeding** for quick testing

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
│   │   ├── recovery.go          # Panic recovery middleware
│   │   └── requestid.go         # Request ID middleware
│   ├── models/
│   │   ├── book.go              # Book model and validation
│   │   ├── book_test.go         # Book model tests
│   │   ├── errors.go            # Domain errors
│   │   ├── filters.go           # Filter models and logic
│   │   ├── filters_test.go      # Filter tests
│   │   └── pagination.go        # Pagination models
│   └── storage/
│       ├── storage.go           # Storage interface
│       ├── memory.go            # In-memory implementation
│       ├── memory_test.go       # Storage tests
│       └── memory_bench_test.go # Performance benchmarks
├── pkg/
│   └── logger/
│       └── logger.go            # Logging utilities
├── api/
│   └── openapi.yaml             # OpenAPI 3.0 specification
├── scripts/
│   └── seed.go                  # Sample data seeder
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
- `GET /books` - Get all books (with pagination and filtering)
  - Query parameters:
    - `page` - Page number (default: 1)
    - `page_size` - Items per page (default: 10, max: 100)
    - `title` - Filter by title (case-insensitive, partial match)
    - `author` - Filter by author (case-insensitive, partial match)
    - `search` - Search in both title and author
- `POST /books` - Create a new book
- `GET /books/{id}` - Get a book by ID
- `PUT /books/{id}` - Update a book (full update)
- `PATCH /books/{id}` - Update a book (partial update)
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
make bench      # Run benchmarks
make seed       # Seed sample data
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

### Get All Books (with Pagination)

```bash
curl http://localhost:8080/books?page=1&page_size=10
```

Response:
```json
{
  "data": [
    {
      "id": 123456,
      "title": "The Go Programming Language",
      "author": "Alan A. A. Donovan"
    }
  ],
  "page": 1,
  "page_size": 10,
  "total": 1,
  "total_pages": 1
}
```

### Search Books

```bash
# Search in both title and author
curl "http://localhost:8080/books?search=Go"

# Filter by title
curl "http://localhost:8080/books?title=Programming"

# Filter by author
curl "http://localhost:8080/books?author=Donovan"

# Combine filters with pagination
curl "http://localhost:8080/books?author=Martin&page=1&page_size=5"
```

### Get a Book by ID

```bash
curl http://localhost:8080/books/123456
```

### Update a Book

```bash
curl -X PUT http://localhost:8080/books/123456 \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Updated Title",
    "author": "Updated Author"
  }'
```

### Delete a Book

```bash
curl -X DELETE http://localhost:8080/books/123456
```

### Seed Sample Data

```bash
# Make sure the server is running first
make seed
```

This will create 20 sample books in the database.

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

Run benchmarks:
```bash
make bench
```

Example benchmark output:
```
BenchmarkMemoryStorage_Create-8          1000000   1024 ns/op   256 B/op   2 allocs/op
BenchmarkMemoryStorage_GetAll-8          5000000    312 ns/op   128 B/op   1 allocs/op
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

## API Documentation

The API is documented using OpenAPI 3.0 specification. You can find the spec at:
- `api/openapi.yaml`

To view the documentation:
1. Install a tool like [Swagger UI](https://swagger.io/tools/swagger-ui/) or [Redoc](https://github.com/Redocly/redoc)
2. Open the `api/openapi.yaml` file

Online viewers:
- https://editor.swagger.io/ (paste the YAML content)

## Performance

The application includes comprehensive benchmarks to ensure optimal performance:

- **Create operations**: ~1000 ns/op
- **Read operations**: ~300 ns/op
- **Concurrent reads**: Highly optimized with RWMutex
- **Thread-safe**: All operations are protected by mutexes

Run `make bench` to see detailed performance metrics.

## Future Enhancements

- [ ] Database integration (PostgreSQL, MongoDB)
- [ ] Authentication and authorization (JWT, OAuth)
- [ ] Rate limiting middleware
- [ ] Caching layer (Redis)
- [ ] Full-text search
- [ ] Sorting options
- [ ] Metrics and monitoring (Prometheus)
- [ ] GraphQL support
- [ ] WebSocket support for real-time updates

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
