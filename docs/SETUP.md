# Setup and Development Guide

## Prerequisites
- Go 1.21 or higher
- PostgreSQL 15 or higher
- Redis 7 or higher
- Docker and Docker Compose (optional, for containerized setup)
- Make (for using Makefile commands)

## Local Development Setup

### 1. Clone the repository
```bash
git clone https://github.com/backrock/saas-backend.git
cd saas-backend
```

### 2. Install dependencies
```bash
go mod download
go mod tidy
```

### 3. Environment setup
```bash
cp .env.example .env
# Edit .env and set your local configuration
```

### 4. Database setup
#### Option A: Using Docker Compose (Recommended)
```bash
make docker-up
```

#### Option B: Manual setup
```bash
# Create database
createdb saas_db

# Run migrations
psql saas_db < migrations/001_initial_schema.sql
```

### 5. Generate code from SQL (if using sqlc)
```bash
make sqlc
```

### 6. Run the application
```bash
make run
```

The server will start on `http://localhost:8080`

## Running Tests
```bash
make test
```

## Code Generation
### Using sqlc
sqlc generates type-safe code from SQL queries:
```bash
make sqlc
```

This generates code in `internal/db/` from queries in `sqlc/queries/`

## Docker Development

### Start all services
```bash
make docker-up
```

### View logs
```bash
make docker-logs
```

### Stop services
```bash
make docker-down
```

### Rebuild containers
```bash
make docker-build
```

## Database Migrations

### Running migrations
```bash
make migrate
```

### Creating new migrations
1. Create a new file: `migrations/002_your_migration_name.sql`
2. Write your SQL statements
3. Run migrations

## Project Structure

```
saas-backend/
├── cmd/
│   └── server/
│       └── main.go              # Application entry point
├── internal/
│   ├── handler/                 # HTTP handlers
│   ├── service/                 # Business logic
│   ├── repository/              # Data access layer
│   ├── model/                   # Data models
│   ├── middleware/              # HTTP middleware
│   └── db/                      # Generated sqlc code
├── migrations/                  # Database migrations
├── sqlc/
│   └── queries/                 # SQL queries for sqlc
├── docker/
│   └── Dockerfile              # Docker image definition
├── docs/                        # Documentation
├── go.mod                       # Go module file
├── go.sum                       # Go dependencies lock file
├── docker-compose.yml          # Docker compose configuration
├── Makefile                    # Development commands
└── README.md                   # Project README
```

## Common Commands

```bash
# Build the application
make build

# Run the application
make run

# Run tests
make test

# Format code
make fmt

# Run linter
make lint

# Clean build artifacts
make clean

# View help
make help
```

## Debugging

### Using VS Code
1. Install the Go extension
2. Create `.vscode/launch.json`:
```json
{
  "version": "0.2.0",
  "configurations": [
    {
      "name": "Connect to Delve",
      "type": "go",
      "request": "launch",
      "mode": "debug",
      "program": "${workspaceFolder}/cmd/server",
      "env": {},
      "args": []
    }
  ]
}
```
3. Press F5 to start debugging

## Performance Optimization

### Database
- Ensure proper indexing
- Use connection pooling (configured in datasource)
- Monitor query performance

### Caching
- Redis is available for caching
- Implement cache warming strategies

### API
- Use pagination for list endpoints
- Implement rate limiting
- Use gzip compression

## Security Checklist

- [ ] Change JWT_SECRET in production
- [ ] Enable HTTPS in production
- [ ] Set up proper CORS policies
- [ ] Use strong database passwords
- [ ] Enable database SSL
- [ ] Implement rate limiting
- [ ] Add input validation
- [ ] Use environment variables for sensitive data
- [ ] Implement logging and monitoring
- [ ] Regular security updates

## Troubleshooting

### Database connection error
```
Check if PostgreSQL is running and configuration in .env is correct
```

### Redis connection error
```
Check if Redis is running and accessible
```

### Port already in use
```bash
# Find the process using the port
lsof -i :8080

# Kill the process
kill -9 <PID>
```

### Docker issues
```bash
# Clean up docker resources
docker system prune -a

# Rebuild containers
make docker-build
```

## Next Steps

1. Implement user registration endpoint
2. Implement user login with JWT
3. Set up Redis caching
4. Add database connection pooling
5. Implement error handling
6. Add comprehensive logging
7. Write unit tests
8. Set up CI/CD pipeline
