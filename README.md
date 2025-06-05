# Project Name

A high-performance web API built with Go Fiber framework.

## 📋 Prerequisites

- Go 1.24 or higher
- Make (for running Makefile commands)

## 🛠 Installation

1. Clone the repository
```bash
git clone https://github.com/tinoimammp/imagyn_go.git
cd imagyn_go
```
2. Install dependencies using make command
```bash
make deps
```
3. Create a `.env` file in the root directory and fill it with your configuration settings using the values from `.example.env`.
```bash
cp .env.example .env
```

## 🚀 Running the Application
1. Run the application
```bash
make run
```
2. Build the application
```bash
make build
```
3. For database migration, we use `migrate`(https://github.com/golang-migrate/migrate). Run `migrate` command:
```bash
migrate -path migrations -database "postgres://db_username:db_password@localhost:5432/db_name?sslmode=disable" up
```

## 🔧 Available Make Commands
- `make run` - Run the application
- `make build` - Build the application
- `make dev` - Build and run the application
- `make test` - Run tests
- `make clean` - Clean build files
- `make deps` - Download and update dependencies

## 🏗 Project Structure
```
├── cmd/                            # Entry point
│   └── main.go
├── internal/                       # Business logic
│   ├── domain/                     # Entity definitions (core business rules)
│   │   ├── example/
│   │   │   └── example.go          # Domain model
│   │   │   └── repository.go       # Repository interface
│   │   │   └── service.go          # Domain service
│   │   │   └── value_objects.go    # Value objects
│   │   └── .../
│   │       └── ...
│   ├── application/            
│   │   └── example/
│   │       ├── usecases/
│   │       │   └── create.go
│   │       └── dto/
│   │           └── example_dto.go
│   ├── infrastructure/
│   │   └── config/                 # Configuration related
│   │       ├── config.go
│   │       └── database.go         
│   │   └── persistence/
│   │       └── postgres/
│   │           ├── provider.go     # Configuration for Postgres
│   │           └── example_repository.go
│   └── interfaces/           
│       └── http/
│           └── example_handler.go 
├── migration/                      # Migration files
│   ├── 001_create_examples.sql
│   └── ...
├── .env.example                    # Environment variables 
├── go.mod
├── Makefile
└── README.md
```