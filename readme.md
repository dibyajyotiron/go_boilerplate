# Go Boilerplate Project

This project is a Go boilerplate with a Domain-Driven Design (DDD) architecture, utilizing Gin for HTTP routing, GORM for ORM, and Consul for service discovery. The project is structured to allow easy modularization and maintainability, so at times, it decouples itself slightly from DDD to make things simpler.

Note: This is work in progress, changes will be made to this repository and the structure as well over time as the need arises for easier project management.

## Project Structure
```
myproject/
├── Dockerfile
├── Dockerfile.test
├── cmd
│ └── main.go
├── docker-compose-test.yml
├── docker-compose.yaml
├── go.mod
├── go.sum
├── internal
│ ├── auth
│ │ └── auth.go
│ ├── config
│ │ ├── config.go
│ │ ├── local.yaml
│ │ ├── prod.yaml
│ │ └── stage.yaml
│ ├── db
│ │ └── db.go
│ ├── discovery
│ │ ├── consul.go
│ │ └── discovery.go
│ ├── queue
│ ├── router
│ │ └── router.go
│ └── user
│ ├── aggregate
│ │ ├── user.go
│ │ └── user_test.go
│ ├── handler
│ │ ├── user_handler.go
│ │ └── user_handler_test.go
│ ├── repository
│ │ ├── user_repository.go
│ │ ├── user_repository_gorm.go
│ │ ├── user_repository_gorm_test.go
│ │ └── user_repository_mock.go
│ └── service
│ ├── user_service.go
│ └── user_service_test.go
└── scripts
└── test.sh
```

## Getting Started

### Prerequisites

- Go 1.20 or later
- Docker
- Docker Compose

### Setup

1. **Clone the repository**

    ```sh
    git clone https://github.com/yourusername/go_boilerplate.git
    cd go_boilerplate
    ```

2. **Environment Variables**

    Create a `.env` file in the root of the project and add the following environment variables:

    ```env
    DATABASE_DSN="host=localhost user=postgres password=yourpassword dbname=yourdb port=5432 sslmode=disable"
    CONSUL_ADDRESS="localhost:8500"
    ```

3. **Configuration Files**

    Ensure you have the following configuration files in the `internal/config` directory:

    - `local.yaml`
    - `stage.yaml`
    - `prod.yaml`

    Example `local.yaml`:

    ```yaml
    database:
      dsn: ""

    consul:
      address: "localhost:8500"
      serviceName: "userService"
      serviceID: "userServiceID"
    ```

### Building and Running

1. **Build the Docker Image**

    ```sh
    docker build -t go_boilerplate .
    ```

2. **Run with Docker Compose**

    ```sh
    docker-compose up
    ```

3. **Access the Application**

    The application will be available at `http://localhost:8080`.

### Running Tests

1. **Build and Run Tests with Docker Compose**

    ```sh
    docker-compose -f docker-compose-test.yml up --abort-on-container-exit --exit-code-from test
    ```

    This command builds the test Docker image and runs the tests, ensuring that the pipeline stops if any tests fail.

### Project Structure Explanation

- **cmd/main.go**: Entry point of the application. It sets up configuration, database, service discovery, and starts the HTTP server.
- **internal/config**: Contains configuration logic and environment-specific configuration files.
- **internal/db**: Database connection logic using GORM.
- **internal/discovery**: Service discovery logic using Consul.
- **internal/user/aggregate**: Domain models for the user.
- **internal/user/handler**: HTTP handlers for user-related endpoints.
- **internal/user/repository**: Data access logic for the user, including GORM and mock implementations.
- **internal/user/service**: Business logic for the user.
- **internal/router**: Router setup using Gin.
- **scripts/test.sh**: Script to run tests.

### Key Concepts

- **Domain-Driven Design (DDD)**: Emphasizes a clear separation between the domain (core business logic) and other aspects of the application (infrastructure, application services, etc.).
- **Gin**: A web framework for Go.
- **GORM**: An ORM library for Go.
- **Consul**: A service discovery and configuration tool.

### Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/your-feature`)
3. Commit your changes (`git commit -m 'Add some feature'`)
4. Push to the branch (`git push origin feature/your-feature`)
5. Create a new Pull Request

### License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

