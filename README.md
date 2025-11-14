# Go ADK Starter Kit

Minimal boilerplate for building AI agents with Google ADK in Go.

## Features

- Sequential blog pipeline (Outline → Writer → Editor)
- Structured logging with `slog` (debug, info, error levels)
- Agent output logger with session state tracking (optional)
- Command-line flags for configuration (log level, agent logger, prompt)
- Clean, minimal codebase

## Project Structure

```
go-adk-starter-kit/
├── cmd/
│   └── go-adk-starter-kit/
│       └── main.go          # Application entry point
├── internal/
│   ├── config/
│   │   └── config.go        # Configuration management
│   ├── logger/
│   │   └── logger.go        # Structured logging with slog
│   └── agents/
│       └── blog/
│           └── blog.go      # Blog pipeline agent
├── go.mod                   # Go module definition
├── go.sum                   # Go module checksums
├── Makefile                 # Build and development tasks
├── .env                     # Environment variables (git ignored)
├── .gitignore               # Git ignore patterns
└── README.md                # This file
```

This follows idiomatic Go project structure:
- `cmd/` - Application entry points (main packages)
- `internal/` - Private application code (not importable by external projects)
- `internal/config/` - Configuration management
- `internal/logger/` - Logging utilities
- `internal/agents/` - Agent implementations

## Setup

1. Clone or copy this starter kit

2. Install dependencies:
```bash
go mod download
```

3. Set up your API key:
```bash
# Create .env file and add your GOOGLE_API_KEY
echo "GOOGLE_API_KEY=your_api_key_here" > .env
```

## Usage

Run with default prompt:
```bash
go run cmd/go-adk-starter-kit/main.go
```

Run with custom prompt:
```bash
go run cmd/go-adk-starter-kit/main.go -prompt "Write a blog post about Golang best practices"
```

Or build and run:
```bash
go build -o bin/go-adk-starter-kit ./cmd/go-adk-starter-kit
./bin/go-adk-starter-kit
```

### Using Makefile

The project includes a Makefile for common tasks:

```bash
# Build the application
make build

# Run with default settings
make run

# Run with debug logging
make run-debug

# Run without agent logger
make run-no-logger

# Clean build artifacts
make clean

# Install dependencies
make deps

# Format code
make fmt

# Show all available targets
make help
```

### Command Line Flags

- `-log-level`: Set logging level (`debug`, `info`, `error`) - Default: `info`
- `-agent-logger`: Enable/disable agent output logging (`true`, `false`) - Default: `true`
- `-prompt`: Blog prompt to process - Default: uses `blog.DefaultPrompt()`

### Examples

Debug mode with agent logger disabled:
```bash
go run cmd/go-adk-starter-kit/main.go -log-level debug -agent-logger=false
```

Error-only logs with custom prompt:
```bash
go run cmd/go-adk-starter-kit/main.go -log-level error -prompt "AI trends in 2024"
```

Full debug with agent outputs:
```bash
go run cmd/go-adk-starter-kit/main.go -log-level debug -agent-logger=true -prompt "Microservices vs Monoliths"
```

## How It Works

The blog pipeline uses a **sequential agent** pattern:

1. **OutlineAgent** - Creates a structured outline
2. **WriterAgent** - Writes content based on the outline
3. **EditorAgent** - Polishes and refines the draft

Each agent:
- Reads input from session state
- Processes the content
- Stores output in session state for the next agent

## Logging

The starter kit uses structured logging with `slog`:

- **Log Levels**: `debug`, `info`, `error` - configurable via `-log-level` flag
- **Application logs**: JSON format to stdout with structured fields
- **Agent outputs**: Pretty-printed to console (can be disabled with `-agent-logger=false`)
- **Session state**: Tracked and logged for debugging
- **Debug mode**: Provides detailed traces of agent execution, session details, and content lengths

## Extending

To add your own agent:

1. Create a new package under `internal/agents/`
2. Implement a `Build()` function that returns `agent.Agent`
3. Define your agent pipeline (sequential, parallel, or loop)
4. Update `cmd/go-adk-starter-kit/main.go` to use your agent
5. Update `internal/logger/` to track your agent's outputs
6. Add configuration options to `internal/config/` if needed

### Adding a New Command

To add additional entry points:

1. Create a new directory under `cmd/` (e.g., `cmd/my-tool/`)
2. Add a `main.go` file with your application logic
3. Import from `internal/` packages as needed
4. Build with: `go build -o bin/my-tool ./cmd/my-tool`

## Dependencies

- `google.golang.org/adk` - Google Agent Development Kit
- `google.golang.org/genai` - Gemini API client
- `github.com/joho/godotenv` - Environment variable management
- `log/slog` - Structured logging (stdlib)

## Contributing

Contributions are welcome! Here's how to get started:

### Development Setup

1. Fork the repository
2. Clone your fork:
```bash
git clone https://github.com/your-username/go-adk-starter-kit.git
cd go-adk-starter-kit
```

3. Install dependencies:
```bash
go mod download
```

4. Create a feature branch:
```bash
git checkout -b feature/your-feature-name
```

### Contribution Guidelines

- **Code Style**: Follow standard Go conventions
  - Run `go fmt ./...` before committing
  - Use meaningful variable and function names
  - Add comments for exported functions and types

- **Testing**: Add tests for new features
  - Place tests in `*_test.go` files
  - Run tests with `go test ./...`

- **Commits**: Write clear commit messages
  - Use present tense ("Add feature" not "Added feature")
  - Reference issues if applicable

- **Documentation**: Update README.md if needed
  - Document new flags or configuration options
  - Add examples for new features

### Submitting Changes

1. Ensure your code passes all checks:
```bash
go fmt ./...
go vet ./...
go test ./...
```

2. Commit your changes:
```bash
git add .
git commit -m "Add: description of your changes"
```

3. Push to your fork:
```bash
git push origin feature/your-feature-name
```

4. Open a Pull Request:
   - Describe what changes you made and why
   - Reference any related issues
   - Wait for review and address feedback

### Ideas for Contributions

- Add new agent patterns (parallel, loop, conditional)
- Implement additional configuration options
- Add support for other LLM providers
- Improve error handling and logging
- Add integration tests
- Create example use cases
- Improve documentation

### Reporting Issues

Found a bug or have a suggestion? [Open an issue](https://github.com/sanjayshr/go-adk-starter-kit/issues) with:
- Clear description of the problem or suggestion
- Steps to reproduce (for bugs)
- Expected vs actual behavior
- Your environment (Go version, OS)

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

