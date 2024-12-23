# Bookstore Management API

## Introduction

The Bookstore Management API is a robust backend service built with Go, Gin, and PostgreSQL. It provides comprehensive functionalities to manage a bookstore, including creating, reading, updating, and deleting book records. This project is designed to be scalable, maintainable, and easily deployable using Docker.

## Project Structure

```plaintext
 ├── cmd
 │ ├── main.go
 │ └── migrate.go
 ├── pkg
 │ ├── config
 │ ├── controllers
 │ ├── database
 │ └── routes
 ├── tests
 ├── Dockerfile
 ├── docker-compose.yml
 ├── go.mod
 └── README.md
```

### Directory Breakdown

#### `cmd`

Contains the entry points for the application.

- **`main.go`**: Initializes configurations, sets up the Gin router, registers routes, and starts the server.
- **`migrate.go`**: Handles database migrations to ensure the database schema is up-to-date.

#### `pkg`

Houses the core packages of the application, each responsible for a specific functionality.

- **`config`**: Manages application and database configurations. It initializes configurations from environment variables and establishes database connections.
  
- **`controllers`**: Contains controller structs and methods that handle HTTP requests and responses. Controllers interact with services to perform business logic.
  
- **`database`**: Includes services that interact directly with the database. It contains the business logic for CRUD operations on book records.
  
- **`routes`**: Defines and registers API routes. It groups related endpoints and associates them with their respective controllers.

#### `tests`

Includes all test files and test-related resources. This directory ensures that the application behaves as expected through unit and integration tests.

#### `Dockerfile`

Defines the Docker image configuration for the application. It specifies the base image, copies the necessary files, installs dependencies, and sets the entry point for the container.

#### `docker-compose.yml`

Facilitates multi-container Docker applications. It defines services, networks, and volumes, enabling easy setup of the application along with its dependencies like the PostgreSQL database.

#### `go.mod`

Specifies the module path and lists all dependencies required for the application. It ensures consistent dependency management and versioning.

#### `README.md`

Provides comprehensive documentation of the project, including setup instructions, API endpoints, and contribution guidelines.

## Libraries Used

### Main Dependencies

- [gin-gonic/gin](https://github.com/gin-gonic/gin) - Web framework
- [lib/pq](https://github.com/lib/pq) - PostgreSQL driver
- [joho/godotenv](https://github.com/joho/godotenv) - Environment configuration
- [gin-contrib/cors](https://github.com/gin-contrib/cors) - CORS middleware

### Testing Dependencies

- [stretchr/testify](https://github.com/stretchr/testify) - Testing assertions
- [DATA-DOG/go-sqlmock](https://github.com/DATA-DOG/go-sqlmock) - SQL mocking
- [stretchr/objx](https://github.com/stretchr/objx) - Mock object generation

## Setup

1. **Clone the repository**

    ```bash
    git clone https://github.com/yourusername/gin-go-PostgresSQL-Bookstore-Management-Api.git
    cd gin-go-PostgresSQL-Bookstore-Management-Api
    ```

2. **Copy environment variables**

    ```bash
    cp .env.example .env
    ```

    - Edit `.env` with your database credentials.

## Migrations

### Development

- **Run Migrations**

    ```bash
    go run cmd/migrate.go
    ```

### Production

- **Using Docker Compose**

    ```bash
    docker-compose up --build
    docker-compose exec app go run cmd/migrate.go
    ```

## Dockerfile Run

- **Build and Run Docker Container**

    ```bash
    docker build -t bookstore-api .
    docker run -p 8080:8080 bookstore-api
    ```

## Running Tests

- **Run All Tests**

    ```bash
    go test ./... -v
    ```

- **Run Tests with Coverage**

    ```bash
    go test ./... -coverprofile=coverage.out
    go tool cover -html=coverage.out
    ```

## API Endpoints

### GET /books

- **Description:** Retrieve a list of all books.
- **Response:**

    ```json
    [
      {
        "id": 1,
        "name": "Test Book",
        "author": "Test Author",
        "publication": "Test Publication"
      }
    ]
    ```

### GET /books/:bookID

- **Description:** Retrieve a single book by its ID.
- **Response:**

    ```json
    {
      "id": 1,
      "name": "Test Book",
      "author": "Test Author",
      "publication": "Test Publication"
    }
    ```

### POST /books

- **Description:** Create a new book.
- **Request:**

    ```json
    {
      "name": "New Book",
      "author": "New Author",
      "publication": "New Publication"
    }
    ```

- **Response:**

    ```json
    {
      "id": 2,
      "name": "New Book",
      "author": "New Author",
      "publication": "New Publication"
    }
    ```

### PUT /books/:bookID

- **Description:** Update an existing book by its ID.
- **Request:**

    ```json
    {
      "name": "Updated Book",
      "author": "Updated Author",
      "publication": "Updated Publication"
    }
    ```

- **Response:**

    ```json
    {
      "id": 1,
      "name": "Updated Book",
      "author": "Updated Author",
      "publication": "Updated Publication"
    }
    ```

### DELETE /books/:bookID

- **Description:** Delete a book by its ID.
- **Response:**

    ```json
    {
      "message": "Book deleted successfully."
    }
    ```

## How to Contribute

1. Fork the repository.
2. Create a new branch:

    ```bash
    git checkout -b feature/YourFeature
    ```

3. Commit your changes:

    ```bash
    git commit -m "Add Your Feature"
    ```

4. Push to the branch:

    ```bash
    git push origin feature/YourFeature
    ```

5. Open a pull request.

## License

This project is licensed under the MIT License. You are free to use, modify, and distribute this software for commercial and personal purposes.
