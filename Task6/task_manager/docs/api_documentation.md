# Task Manager API Documentation

## Project Overview

Task Manager is a RESTful API built with Go for managing tasks. It allows users to create, view, update, and delete tasks. The project is organized using MVC principles for clarity and maintainability.

- **Language:** Go
- **Main Features:**
  - CRUD operations for tasks
  - JSON-based API
  - Simple error handling
- **Project Structure:**
  - `main.go`: Entry point
  - `controllers/`: Handles HTTP requests
  - `models/`: Data models
  - `data/`: Business logic/service layer
  - `router/`: API routing
  - `docs/`: Documentation

## Getting Started

### Prerequisites

- Go 1.18 or higher installed
- A running MongoDB instance (local or Atlas)

### MongoDB Configuration

1. **Set up MongoDB:**
   - For local development, install MongoDB Community Edition from [mongodb.com](https://www.mongodb.com/try/download/community).
   - For cloud, create a free cluster on [MongoDB Atlas](https://www.mongodb.com/atlas/database).

2. **Configure connection string:**
   - Create a `.env` file in the project root.
   - Add your MongoDB connection string:
     ```env
     MONGO_API_URL=mongodb://localhost:27017
     # Or for Atlas:
     MONGO_API_URL=mongodb+srv://<username>:<password>@<cluster-url>/test?retryWrites=true&w=majority
     ```
   - For Atlas, ensure your IP is whitelisted in Network Access and your cluster is running.

3. **Verify connection:**
   - The server will print a message when successfully connected to MongoDB.

### Installation & Running

1. Clone the repository:
   ```sh
   git clone https://github.com/Kad-19/a2sv-backend-project-phase
   cd Task4/task_manager
   ```
2. Download dependencies:
   ```sh
   go mod tidy
   ```
3. Run the server:
   ```sh
   go run main.go
   ```
   The API will be available at `http://localhost:8080`.

### Testing the API

You can use tools like [Postman](https://www.postman.com/) or [curl](https://curl.se/) to test the endpoints.

#### Example curl requests:

- **Create a task:**
  ```sh
  curl -X POST http://localhost:8080/tasks -H "Content-Type: application/json" -d '{"title":"Test","description":"Try out the API","dueDate": "2025-07-15T15:23:27.0682587+03:00"}'
  ```
- **Get all tasks:**
  ```sh
  curl http://localhost:8080/tasks
  ```
- **Get a task by ID:**
  ```sh
  curl http://localhost:8080/tasks/1
  ```
- **Update a task:**
  ```sh
  curl -X PUT http://localhost:8080/tasks/1 -H "Content-Type: application/json" -d '{"title":"Updated","description":"Updated desc","dueDate": "2025-07-15T15:23:27.0682587+03:00"}'
  ```
- **Delete a task:**
  ```sh
  curl -X DELETE http://localhost:8080/tasks/1
  ```

---

## Endpoints

### 1. Create a Task

- **URL:** `/tasks`
- **Method:** `POST`
- **Request Body:**

```json
{
  "title": "string",
  "description": "string",
  "dueDate": "string",
}
```

- **Response:**

```json
{
  "id": 1,
  "title": "string",
  "description": "string",
  "dueDate": "string",
  "Status": "string"
}
```

- **Status Codes:**
  - `201 Created`: Task created successfully
  - `400 Bad Request`: Invalid input

---

### 2. Get All Tasks

- **URL:** `/tasks`
- **Method:** `GET`
- **Response:**

```json
[
  {
    "id": 1,
    "title": "string",
    "description": "string",
    "dueDate": "string",
    "Status": "string"
  }
  // ...more tasks
]
```

- **Status Codes:**
  - `200 OK`: Success

---

### 3. Get Task by ID

- **URL:** `/tasks/{id}`
- **Method:** `GET`
- **Response:**

```json
{
  "id": 1,
  "title": "string",
  "description": "string",
  "dueDate": "string",
  "Status": "string"
}
```

- **Status Codes:**
  - `200 OK`: Success
  - `404 Not Found`: Task not found

---

### 4. Update Task

- **URL:** `/tasks/{id}`
- **Method:** `PUT`
- **Request Body:**

```json
{
  "title": "string",
  "description": "string",
  "dueDate": "string",
  "Status": "string"
}
```

- **Response:**

```json
{
  "id": 1,
  "title": "string",
  "description": "string",
  "dueDate": "string",
  "Status": "string"
}
```

- **Status Codes:**
  - `200 OK`: Task updated
  - `400 Bad Request`: Invalid input
  - `404 Not Found`: Task not found

---

### 5. Delete Task

- **URL:** `/tasks/{id}`
- **Method:** `DELETE`
- **Response:**

```json
{
  "message": "Task deleted successfully"
}
```

- **Status Codes:**
  - `200 OK`: Task deleted
  - `404 Not Found`: Task not found

---

## Error Response Format

```json
{
  "error": "Error message"
}
```

## Model

- **Task Object:**

```json
{
  "id": 1,
  "title": "string",
  "description": "string",
  "dueDate": "string",
  "Status": "string"
}
```

## Notes

- All endpoints accept and return JSON.
- Ensure to set `Content-Type: application/json` in requests.
