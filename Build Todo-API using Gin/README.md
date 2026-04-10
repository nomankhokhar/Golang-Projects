# Todo API using Gin

A simple yet elegant RESTful API for managing todo items built with Go and the Gin web framework.

## Features

- **CRUD Operations**: Create, Read, Update, and Delete todo items
- **RESTful API**: Clean and intuitive API endpoints
- **In-Memory Storage**: Fast operations with seed data
- **Type-Safe**: Built with Go's strong typing system
- **Easy to Use**: Simple controller and route structure

## Project Structure

```
├── main.go                 # Entry point of the application
├── go.mod                  # Go module definition
├── go.sum                  # Module checksums
├── controllers/
│   └── todo_controller.go  # Handler logic for todo operations
├── models/
│   └── todo.go            # Data models and store logic
└── routes/
    └── routes.go          # API routes definition
```

## Getting Started

### Prerequisites

- Go 1.25.0 or higher
- Gin Web Framework (`github.com/gin-gonic/gin`)

### Installation

1. Clone or navigate to the project directory
2. Install dependencies:
```bash
go mod download
```

3. Run the application:
```bash
go run main.go
```

The API will start on `http://localhost:8080`

## API Endpoints

### Base URL
```
http://localhost:8080/api/v1
```

### Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/todos` | Retrieve all todos |
| GET | `/todos/:id` | Retrieve a specific todo by ID |
| POST | `/todos` | Create a new todo |
| PUT | `/todos/:id` | Update an existing todo |
| DELETE | `/todos/:id` | Delete a todo |

## Usage Examples

### Get All Todos
```bash
curl http://localhost:8080/api/v1/todos
```

### Get Todo by ID
```bash
curl http://localhost:8080/api/v1/todos/1
```

### Create a New Todo
```bash
curl -X POST http://localhost:8080/api/v1/todos \
  -H "Content-Type: application/json" \
  -d '{"title":"New Task","status":"Incomplete"}'
```

### Update a Todo
```bash
curl -X PUT http://localhost:8080/api/v1/todos/1 \
  -H "Content-Type: application/json" \
  -d '{"title":"Updated Task","status":"Complete"}'
```

### Delete a Todo
```bash
curl -X DELETE http://localhost:8080/api/v1/todos/1
```

## Data Model

### Todo
```go
type Todo struct {
    ID     int    `json:"id"`      // Unique identifier
    Title  string `json:"title"`   // Task description
    Status string `json:"status"`  // Task status
}
```

## Seed Data

The application starts with the following pre-loaded todos:
- ID: 1, Title: "Learn Golang", Status: "Incomplete"
- ID: 2, Title: "Build API", Status: "Incomplete"

## Architecture

### Controllers (`controllers/todo_controller.go`)
- `TodoController`: Handles all HTTP request logic
- `TodoHandler`: Interface defining controller behavior
- Request validation and error handling

### Models (`models/todo.go`)
- `Todo`: Data structure for todo items
- `TodoStore`: In-memory storage managing todo data
- CRUD operations on the todo collection

### Routes (`routes/routes.go`)
- `SetupRouter()`: Initializes Gin engine and registers all routes
- Grouping API endpoints under `/api/v1` namespace

## Key Components

### TodoController
Implements request handling with the following methods:
- `GetAllTodos()`: Returns all todos
- `GetTodoByID()`: Returns a single todo
- `CreateTodo()`: Creates a new todo with auto-generated ID
- `UpdateTodo()`: Updates an existing todo
- `DeleteTodo()`: Removes a todo

### TodoStore
Manages in-memory todo storage:
- Auto-incrementing ID generation
- Search by ID functionality
- CRUD operations

## Technical Details

- **Framework**: Gin Gonic (high-performance HTTP web framework)
- **Language**: Go 1.25.0
- **Storage**: In-memory (no database)
- **Server Port**: 8080

## Future Enhancements

- Database integration (MongoDB, PostgreSQL)
- Authentication and authorization
- Request validation and error handling improvements
- Pagination for todo lists
- Filtering and sorting capabilities
- Persistent storage with migrations
- Unit and integration tests
- Docker containerization

## License

This project is part of the Golang-Projects collection.
