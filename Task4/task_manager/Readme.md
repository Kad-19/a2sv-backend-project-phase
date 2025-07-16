# Task Manager

A RESTful API for managing tasks, built with Go.

## Features

- Create, read, update, and delete tasks
- JSON-based API
- Simple error handling

## Project Structure

- `main.go`: Entry point
- `controllers/`: Handles HTTP requests
- `models/`: Data models
- `data/`: Business logic/service layer
- `router/`: API routing
- `docs/`: Documentation

## API Documentation

- Full API documentation is available in [`docs/api_documentation.md`](docs/api_documentation.md).
- You can also view and test the API using Postman:
  - [Postman API Documentation](https://documenter.getpostman.com/view/33603724/2sB34iiyXp)

## Getting Started

### Prerequisites

- Go 1.18 or higher

### Installation & Running

1. Clone the repository:
   ```sh
   git clone <repo-url>
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

- Use [Postman](https://www.postman.com/) or [curl](https://curl.se/) to test endpoints.
- Example requests are provided in the API documentation file.
