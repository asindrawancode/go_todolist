# Todo List Application

This is a simple Go-based ToDo List application that allows users to register, login, and manage their to-do tasks. It employs the Gorilla Mux for HTTP routing, Gorm for ORM, JSON Web Tokens (JWT) for authentication, and Docker for containerization.

## Features

- User Registration and Authentication
- JWT-based secure access to APIs
- Task creation and retrieval
- Containerized deployment with Docker

## Technologies Used

- Go 1.20
- Gorilla Mux
- Gorm ORM with SQLite
- JWT for authentication
- Docker & Docker Compose

## Getting Started

These instructions will help you set up and run the application locally using Docker.

### Prerequisites

- [Docker](<https://www.docker.com/>)
- [Docker Compose](<https://docs.docker.com/compose/>)

### Installing

1. **Clone the repository**:
   ```bash
   git clone <https://github.com/yourusername/todolist.git>
   cd todolist


1. **Build and start the Docker container**:

    ```bash
    docker-compose up --build
    
    ```

   This command will pull the necessary Docker images, build your application, and start the container(s).

2. **Access the API**:
   The API is now accessible at `http://localhost:8000`.

- **Register a new user**:

    ```bash
    curl -X POST -H "Content-Type: application/json" -d '{"username":"user1","password":"password"}' http://localhost:8000/register
    
    ```

- **Login**:

    ```bash
    curl -X POST -H "Content-Type: application/json" -d '{"username":"user1","password":"password"}' http://localhost:8000/login
    
    ```

  The login response includes a JWT token. Use this token for accessing secured endpoints by attaching it as a Bearer token in the `Authorization` header.

- **Add a To-Do**:

    ```bash
    curl -X POST -H "Authorization: Bearer <token>" -H "Content-Type: application/json" -d '{"task":"Buy groceries"}' http://localhost:8000/api/todolist
    
    ```

- **Retrieve To-Do List**:

    ```bash
    curl -X GET -H "Authorization: Bearer <token>" http://localhost:8000/api/todolist
    
    ```

### Environment Variables

Adjust these settings in `docker-compose.yml`:

- `JWT_SECRET`: The secret key for JWT signing.

### Logging

All API requests are logged to the console with basic request information (method, URI, remote address).

### Persistence

The application uses SQLite to persist data, stored in a Docker volume for data persistence across container restarts.