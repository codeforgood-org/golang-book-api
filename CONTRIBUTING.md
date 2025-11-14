# Contributing to Book API

Thank you for your interest in contributing to the Book API project! This document provides guidelines and instructions for contributing.

## Code of Conduct

By participating in this project, you agree to maintain a respectful and inclusive environment for everyone.

## How to Contribute

### Reporting Bugs

If you find a bug, please create an issue with:
- A clear, descriptive title
- Steps to reproduce the issue
- Expected behavior
- Actual behavior
- Your environment (OS, Go version, etc.)

### Suggesting Enhancements

Enhancement suggestions are welcome! Please create an issue with:
- A clear, descriptive title
- Detailed description of the proposed feature
- Rationale for why this enhancement would be useful
- Possible implementation approach (optional)

### Pull Requests

1. **Fork the repository** and create your branch from `main`
   ```bash
   git checkout -b feature/amazing-feature
   ```

2. **Make your changes**
   - Follow the project's coding style
   - Write clear, concise commit messages
   - Add tests for new functionality
   - Update documentation as needed

3. **Test your changes**
   ```bash
   make test
   make lint
   ```

4. **Commit your changes**
   ```bash
   git commit -m "Add amazing feature"
   ```

5. **Push to your fork**
   ```bash
   git push origin feature/amazing-feature
   ```

6. **Open a Pull Request**
   - Provide a clear description of the changes
   - Reference any related issues
   - Ensure all tests pass
   - Wait for review

## Development Setup

### Prerequisites

- Go 1.21 or higher
- Git
- Make (optional but recommended)

### Setup Steps

1. Clone your fork:
   ```bash
   git clone https://github.com/YOUR_USERNAME/golang-book-api.git
   cd golang-book-api
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Run tests:
   ```bash
   make test
   ```

4. Run the application:
   ```bash
   make run
   ```

## Coding Standards

### Go Style Guide

- Follow the official [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- Use `gofmt` to format your code
- Run `go vet` to check for common mistakes
- Use meaningful variable and function names

### Project Structure

- **`cmd/`** - Application entry points
- **`internal/`** - Private application code
- **`pkg/`** - Public libraries
- Keep packages focused and cohesive
- Use interfaces for dependencies

### Testing

- Write unit tests for all new functionality
- Aim for high test coverage
- Use table-driven tests where appropriate
- Mock external dependencies

Example test:
```go
func TestBookValidate(t *testing.T) {
    tests := []struct {
        name    string
        book    Book
        wantErr bool
    }{
        // test cases
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // test implementation
        })
    }
}
```

### Documentation

- Add godoc comments for public functions and types
- Update README.md for significant changes
- Include examples for complex functionality

### Commit Messages

Write clear commit messages:
- Use present tense ("Add feature" not "Added feature")
- Keep the first line under 50 characters
- Add detailed description after a blank line if needed

Good examples:
```
Add book validation logic

Implement validation for book title and author fields.
Returns appropriate error messages for invalid input.
```

## Testing Guidelines

### Running Tests

```bash
# Run all tests
make test

# Run tests with coverage
make test-cover

# Run specific package tests
go test ./internal/storage/...
```

### Writing Tests

- Test public APIs
- Test edge cases and error conditions
- Use descriptive test names
- Keep tests simple and focused

## Review Process

1. All pull requests require review before merging
2. Address reviewer feedback
3. Keep discussions professional and constructive
4. CI checks must pass

## Questions?

If you have questions about contributing, feel free to:
- Open an issue for discussion
- Reach out to the maintainers

Thank you for contributing!
