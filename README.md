# Task Management System

## Overview

This project is a **Task Management System** built using a microservices architecture in Go. It allows users to register, authenticate, and manage tasks, with notifications for task events. The system is composed of three main services:

- **User Service**: Handles user registration, authentication, and user CRUD operations.
- **Task Service**: Manages tasks (CRUD, filtering, status updates).
- **Notification Service**: Sends notifications for task events (e.g., reminders).

---

## Intution
- Thinking of making the server microservice architecture because of multiple different components which can work independently of one another.
- Initally on high level we have 2 main services. one is TaskService and other is UserService.
- Other Services like authentication, notification, analytics depend on them and can be build later.
- Task Service should handle tasks of all users. It should have all CRUD operation and should send events related to task for other services.
- User Service is basically a basic service to register and deregister users using the services.
- Inter service communications should be communicating via gRPC.
- Other services that don't directly interfere with tasks and users data and need to just read and do some operation on those data should be event driven services.
- events can be emmited through kafka which allows multiple services to work in async manner.

## Design Decisions

- TaskService will not be exposed externally. It will only have grpc and pub sub communication with other services.
- Similarly User service will only have grpc communication with other services and not exposed outside.
- Both these services will communicate with api gateway service which will implement authentication, rate limits etc.
- Api Gateway will not be implemented currently and User service itself will contain these functionalities due to following reason:
  1. For these external database server and other functionalities needs to be implemented which will be time consuming.
  2. Lot of time constraint
  3. Without a client server proper authentication will still be a simple authentication.
- Because of these reasons, User service will have authentication, all the api endpoints and user data CRUD operations.
- A simple notification service which will capture events emitted from task-service through kafka.

---

## How This Project Demonstrates Microservices Concepts

- **Service Decomposition**: Each business capability (user, task, notification) is a separate service.
- **Independent Deployability**: Each service has its own Dockerfile and can be deployed/scaled independently.
- **Inter-Service Communication**: Uses gRPC for efficient, language-agnostic communication and kafka for async communication between services.
- **API Gateway Pattern**: User Service exposes a unified REST API, hiding internal service boundaries.
- **Centralized Authentication**: JWT tokens are issued and validated by the User Service, but used across all endpoints.
- **Resilience**: Each service can fail or restart independently without bringing down the whole system.

---

## Project Structure

```
task-managment-system/
  ├── docker-compose.yml
  ├── notification-service/
  ├── proto/
  ├── task-service/
  ├── user-service/
  └── README.md
```

---

## Running the Service

### Prerequisites

- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/)
- (Optional) [Go](https://golang.org/) for local development

### Quick Start

1. **Clone the repository:**
   ```sh
   git clone https://github.com/yourusername/task-managment-system.git
   cd task-managment-system
   ```

2. **Start all services:**
   ```sh
   make all
   ```

   This will start:
   - User Service (REST API + gRPC)
   - Task Service (gRPC)
   - Notification Service (gRPC)

3. **Access the API:**
   - User Service REST API: `http://localhost:8082`

4. **Stop the services:**
   ```sh
   make kill
   ```

---

## API Documentation

### Authentication

- **Register:** `POST /user/register`
- **Login:** `POST /user/login`
- **All other endpoints require**: `Authorization: Bearer <token>`

### User Endpoints

| Method | Path           | Description         | Auth Required |
|--------|----------------|---------------------|--------------|
| POST   | /user/register | Register new user   | No           |
| POST   | /user/login    | Login, get JWT      | No           |
| GET    | /user          | Get user by token   | Yes          |
| POST    | /user/update   | Update user         | Yes          |
| DELETE | /user          | Delete user         | Yes          |

## Detailed User Endpoints Documentation

### 1. Register User
- **POST** `/user/register`
- **Request Body:**
  ```json
  {
    "name": "John Doe",
    "email": "john@example.com",
    "password": "password123"
  }
  ```
- **Response:**
  ```json
  {
    "id": 1,
    "name": "John Doe",
    "email": "john@example.com"
  }
  ```
- **cURL:**
  ```sh
  curl -X POST http://localhost:8082/user/register \
    -H "Content-Type: application/json" \
    -d '{"name":"John Doe","email":"john@example.com","password":"password123"}'
  ```
- **Postman:**  
  - Method: POST  
  - URL: http://localhost:8082/user/register  
  - Body: raw JSON (as above)

---

### 2. Login
- **POST** `/user/login`
- **Request Body:**
  ```json
  {
    "email": "john@example.com",
    "password": "password123"
  }
  ```
- **Response:**
  ```json
  {
    "token": "mock-jwt-token",
    "user": {
      "id": 1,
      "name": "John Doe",
      "email": "john@example.com"
    }
  }
  ```
- **cURL:**
  ```sh
  curl -X POST http://localhost:8082/user/login \
    -H "Content-Type: application/json" \
    -d '{"email":"john@example.com","password":"password123"}'
  ```
- **Postman:**  
  - Method: POST  
  - URL: http://localhost:8082/user/login  
  - Body: raw JSON (as above)

---

### 3. Get User by ID
- **GET** `/user`
- **Headers:** `Authorization: Bearer user-1`
- **Response:**
  ```json
  {
    "id": 1,
    "name": "John Doe",
    "email": "john@example.com"
  }
  ```
- **cURL:**
  ```sh
  curl -X GET http://localhost:8082/user \
    -H "Authorization: Bearer user-1"
  ```
- **Postman:**  
  - Method: GET  
  - URL: http://localhost:8082/user  
  - Headers: Authorization: Bearer user-1

---

### 4. Update User
- **POST** `/user/update`
- **Headers:** `Authorization: Bearer user-1`
- **Request Body:**
  ```json
  {
    "name": "Jane Doe",
    "email": "jane@example.com",
    "password": "newpasswordhash"
  }
  ```
- **Response:**
  ```json
  {
    "id": 1,
    "name": "Jane Doe",
    "email": "jane@example.com"
  }
  ```
- **cURL:**
  ```sh
  curl -X POST http://localhost:8082/user/update \
    -H "Authorization: Bearer user-1" \
    -H "Content-Type: application/json" \
    -d '{"name":"Jane Doe","email":"jane@example.com","password":"newpasswordhash"}'
  ```
- **Postman:**  
  - Method: POST  
  - URL: http://localhost:8082/user/update 
  - Headers: Authorization: Bearer user-1  
  - Body: raw JSON (as above)

---

### 5. Delete User
- **DELETE** `/user`
- **Headers:** `Authorization: Bearer user-1`
- **Response:** HTTP 204 No Content
- **cURL:**
  ```sh
  curl -X DELETE http://localhost:8082/user \
    -H "Authorization: Bearer user-1"
  ```
- **Postman:**  
  - Method: DELETE  
  - URL: http://localhost:8082/user  
  - Headers: Authorization: Bearer user-1

---

### Task Endpoints

| Method | Path                        | Description           | Auth Required |
|--------|---------------------------- |-----------------------|--------------|
| GET    | /task/list_all_tasks        | List tasks            | Yes          |
| POST   | /task/create                | Create new task       | Yes          |
| GET    | /task?id={id}               | Get task by ID        | Yes          |
| POST   | /task/update?id={id}        | Update task by ID     | Yes          |
| DELETE | /task?id={id}               | Delete task by ID     | Yes          |

## Detailed Task Endpoints Documentation

### 1. List Tasks
- **GET** `/task/list_all_tasks?status=TASK_STATUS_PENDING&page=1&limit=10&sort=date`
- **Headers:** `Authorization: Bearer user-1`
- **Response:**
  ```json
  [
    {
      "id": 1,
      "title": "Task 1",
      "description": "Description",
      "userId": 1,
      "status": "TASK_STATUS_PENDING",
      "createdAt": "...",
      "updatedAt": "...",
      "triggerAt": "...",
      "triggered": false
    }
  ]
  ```
- **cURL:**
  ```sh
  curl -X GET "http://localhost:8082/task/list_all_tasks?status=TASK_STATUS_PENDING&page=1&limit=10&sort=date" \
    -H "Authorization: Bearer user-1"
  ```
- **Postman:**  
  - Method: GET  
  - URL: http://localhost:8082/task/list_all_tasks?status=TASK_STATUS_PENDING&page=1&limit=10&sort=date  
  - Headers: Authorization: Bearer user-1

---

### 2. Create Task
- **POST** `/task/create`
- **Headers:** `Authorization: Bearer user-1`
- **Request Body:**
  ```json
  {
    "title": "New Task",
    "description": "Task details",
    "trigger_at": "2024-06-01T10:00:00Z"
  }
  ```
- **Response:**
  ```json
  {
    "id": 2,
    "title": "New Task",
    "description": "Task details",
    "userId": 1,
    "status": "TASK_STATUS_NO_STATUS",
    "createdAt": "...",
    "updatedAt": "...",
    "triggerAt": "2024-06-01T10:00:00Z",
    "triggered": false
  }
  ```
- **cURL:**
  ```sh
  curl -X POST http://localhost:8082/task/create \
    -H "Authorization: Bearer user-1" \
    -H "Content-Type: application/json" \
    -d '{"title":"New Task","description":"Task details","trigger_at":"2024-06-01T10:00:00Z"}'
  ```
- **Postman:**  
  - Method: POST  
  - URL: http://localhost:8082/task/create  
  - Headers: Authorization: Bearer user-1  
  - Body: raw JSON (as above)

---

### 3. Get Task by ID
- **GET** `/task?id=2`
- **Headers:** `Authorization: Bearer user-1`
- **Response:**
  ```json
  {
    "id": 2,
    "title": "New Task",
    "description": "Task details",
    "userId": 1,
    "status": "TASK_STATUS_NO_STATUS",
    "createdAt": "...",
    "updatedAt": "...",
    "triggerAt": "2024-06-01T10:00:00Z",
    "triggered": false
  }
  ```
- **cURL:**
  ```sh
  curl -X GET "http://localhost:8082/task?id=2" \
    -H "Authorization: Bearer user-1"
  ```
- **Postman:**  
  - Method: GET  
  - URL: http://localhost:8082/task?id=2  
  - Headers: Authorization: Bearer user-1

---

### 4. Update Task
- **POST** `/task/update?id=2`
- **Headers:** `Authorization: Bearer user-1`
- **Request Body:**
  ```json
  {
    "title": "Updated Task",
    "description": "Updated details",
    "status": "TASK_STATUS_COMPLETED",
    "trigger_at": "2024-06-01T12:00:00Z"
  }
  ```
- **Response:** HTTP 200 OK
- **cURL:**
  ```sh
  curl -X POST "http://localhost:8082/task/update?id=2" \
    -H "Authorization: Bearer user-1" \
    -H "Content-Type: application/json" \
    -d '{"title":"Updated Task","description":"Updated details","status":"TASK_STATUS_COMPLETED","trigger_at":"2024-06-01T12:00:00Z"}'
  ```
- **Postman:**  
  - Method: POST  
  - URL: http://localhost:8082/task/update?id=2  
  - Headers: Authorization: Bearer user-1  
  - Body: raw JSON (as above)

---

### 5. Delete Task
- **DELETE** `/task?id=2`
- **Headers:** `Authorization: Bearer user-1`
- **Response:** HTTP 200 OK
- **cURL:**
  ```sh
  curl -X DELETE "http://localhost:8082/task?id=2" \
    -H "Authorization: Bearer user-1"
  ```
- **Postman:**  
  - Method: DELETE  
  - URL: http://localhost:8082/task?id=2  
  - Headers: Authorization: Bearer user-1

---

### Notification Service

- **Internal gRPC service**: Sends notifications for task events (e.g., reminders).
- Not directly exposed to clients; triggered by Task Service.

---

## Development

### Running Locally (without Docker)

1. Start each service in its directory:
   ```sh
   cd user-service && go run cmd/main.go
   cd task-service && go run cmd/main.go
   cd notification-service && go run cmd/main.go
   ```
2. Ensure ports do not conflict (see each service's config).


## Extending the System

- **Add new services**: Define proto contracts, implement service, add to `docker-compose.yml`.
- **Add new endpoints**: Update proto files and REST handlers.
- **Change data storage**: Swap out the storage layer in each service independently.
- **Api Gateway Addition**: Api Gateway service can be easily implemented once database server is created. User service is written such that api gateway functionalities are easily decoupled from user-service.
---

## Contact

For questions or contributions, please open an issue or pull request.

---

**Note:** This project is for educational/demo purposes and may use mock tokens or in-memory storage. For production, add persistent storage, secure JWT handling, and production-ready configurations.
