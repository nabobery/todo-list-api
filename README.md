# Todo List API

https://roadmap.sh/projects/todo-list-api

This project is a RESTful API for managing to-do items. Users can register, log in, and perform CRUD operations on their personal to-do lists. The API uses JWT for secure authentication, MongoDB for data persistence, and comes with auto-generated Swagger documentation.

## Goals

The purpose of this project is to:

- Understand and implement RESTful API design with proper HTTP methods.
- Learn how to create, read, update, and delete resources.
- Implement secure user authentication using JWT.
- Explore MongoDB for data persistence.
- Utilize Swagger for API documentation.
- Handle errors gracefully and apply proper input validation.
- Implement pagination for listing to-do items.

## Features

- **User Authentication:**

  - **Register:** `POST /register` - Create a new user with secure password hashing.
  - **Login:** `POST /login` - Authenticate a user and generate a JWT token.

- **To-Do Operations:**
  - **Create To-do:** `POST /todos` - Add a new to-do item (requires JWT).
  - **Update To-do:** `PUT /todos/{id}` - Modify an existing to-do item (only by its creator).
  - **Delete To-do:** `DELETE /todos/{id}` - Remove a to-do item (only by its creator).
  - **Get To-dos:** `GET /todos?page=1&limit=10` - Retrieve a paginated list of to-do items (requires JWT).

## Technologies Used

- **Language:** Go
- **Framework:** Gin
- **Database:** MongoDB
- **Authentication:** JWT (via `github.com/golang-jwt/jwt/v5`)
- **Documentation:** Swagger (via swaggo)
- **Configuration:** Environment variables loaded via `godotenv`

## Project Flow

## Project flow

[![](https://mermaid.ink/img/pako:eNptkDFuwzAMRa8icI4uoCFADQ1NETRB7S6BFtZWHSG1qIrSEAS5e2kYnlxN5P_vEyIf0NPgwQD73-pj723AMePkopKXMJfQh4SxqE_2WSGrjrQldQws0mGLvZwPCyWQ1FvANrNvseAXsndxIebper-XhFFtGKOuaWscadThn8CpXOVvlHixRBPHNka9dt1ZfcybcVk82-g19dae3sXkRJE97GDyecIwyC0eM-tApk7egZFywHxz4OJTOKyF2nvswZRc_Q4y1fG6NjUNWNYjgvnGHxZVFr8Qrf3zDwawd6M?type=png)](https://mermaid.live/edit#pako:eNptkDFuwzAMRa8icI4uoCFADQ1NETRB7S6BFtZWHSG1qIrSEAS5e2kYnlxN5P_vEyIf0NPgwQD73-pj723AMePkopKXMJfQh4SxqE_2WSGrjrQldQws0mGLvZwPCyWQ1FvANrNvseAXsndxIebper-XhFFtGKOuaWscadThn8CpXOVvlHixRBPHNka9dt1ZfcybcVk82-g19dae3sXkRJE97GDyecIwyC0eM-tApk7egZFywHxz4OJTOKyF2nvswZRc_Q4y1fG6NjUNWNYjgvnGHxZVFr8Qrf3zDwawd6M)

## Project Structure

```bash
todo-list-api/
├── cmd/
│   └── main.go               # Entry point: load config, setup routes, start server
├── config/
│   └── config.go             # Loads environment variables and connects to MongoDB
├── controllers/
│   ├── auth_controller.go    # HTTP handlers for user registration and login
│   └── todo_controller.go    # HTTP handlers for CRUD operations on to-do items
├── docs/                     # Auto-generated Swagger docs (swag init)
├── middlewares/
│   └── auth_middleware.go    # JWT authentication middleware protecting endpoints
├── models/
│   ├── user.go               # User model
│   └── todo.go               # To-do item model
├── repository/
│   ├── user_repository.go    # Data access layer for users in MongoDB
│   └── todo_repository.go    # Data access layer for to-do items in MongoDB
├── routes/
│   └── routes.go             # Registers all routes and attaches controllers and middleware
├── services/
│   ├── auth_service.go       # Business logic for user authentication
│   └── todo_service.go       # Business logic for to-do operations
├── go.mod                    # Module definition file
└── go.sum
```

## Endpoints

### User Authentication

**Register a New User**
`POST /register`
_Request:_

```json
{
  "name": "John Doe",
  "email": "john@doe.com",
  "password": "password"
}
```

_Response:_

```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

**Login User**
`POST /login`
_Request:_

```json
{
  "email": "john@doe.com",
  "password": "password"
}
```

_Response:_

```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

### To-Do Operations

**Create a To-Do Item**
`POST /todos`
_Headers:_ `Authorization: Bearer <token>`
_Request:_

```json
{
  "title": "Buy groceries",
  "description": "Buy milk, eggs, and bread"
}
```

_Response:_

```json
{
  "id": "60d21bae3f1a2c001c8f3c90",
  "title": "Buy groceries",
  "description": "Buy milk, eggs, and bread",
  "user_id": "60d21bae3f1a2c001c8f3c89",
  "created_at": "2023-10-01T12:34:56Z",
  "updated_at": "2023-10-01T12:34:56Z"
}
```

**Update a To-Do Item**
`PUT /todos/{id}`
_Headers:_ `Authorization: Bearer <token>`
_Request:_

```json
{
  "title": "Buy groceries",
  "description": "Buy milk, eggs, bread, and cheese"
}
```

_Response:_

```json
{
  "id": "60d21bae3f1a2c001c8f3c90",
  "title": "Buy groceries",
  "description": "Buy milk, eggs, bread, and cheese",
  "user_id": "60d21bae3f1a2c001c8f3c89",
  "created_at": "2023-10-01T12:34:56Z",
  "updated_at": "2023-10-01T13:00:00Z"
}
```

**Delete a To-Do Item**
`DELETE /todos/{id}`
_Headers:_ `Authorization: Bearer <token>`
_Response:_ HTTP status code `204 No Content`

**Get To-Do Items**
`GET /todos?page=1&limit=10`
_Headers:_ `Authorization: Bearer <token>`
_Response:_

```json
{
  "data": [
    {
      "id": "60d21bae3f1a2c001c8f3c90",
      "title": "Buy groceries",
      "description": "Buy milk, eggs, and bread"
    }
  ],
  "page": 1,
  "limit": 10,
  "total": 1
}
```

## Environment Variables

Create a `.env` file in the root directory to configure the application:

```env
# MongoDB connection string
MONGO_URI="mongodb://localhost:27017"

# MongoDB database name
MONGO_DB="tododb"

# JWT secret key for signing tokens
JWT_SECRET="your_secret_key"

# Port for the API server
PORT="8080"
```

## Installation & Setup

1. **Clone the Repository**

   ```bash
   git clone https://github.com/nabobery/todo-list-api.git
   cd todo-list-api
   ```

2. **Install Dependencies**

   ```bash
   go mod tidy
   ```

3. **Configure Environment**
   Create a `.env` file in the project root with your MongoDB settings, JWT secret, and desired port (see above).

4. **Run the Application**

   ```bash
   go run cmd/main.go
   ```

   The server should start on the port specified (default is `8080`).

## API Documentation

Swagger documentation is auto-generated using swaggo. To generate and view the docs:

1. **Generate Swagger Docs**

   ```bash
   swag init -g cmd/main.go
   ```

2. **Serving Swagger**
   The docs are served at `/swagger/index.html`.

## Testing the API

You may use cURL, Postman, or any other API testing tool. For example, to register a new user:

```bash
curl --location --request POST 'http://localhost:8080/register' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "John Doe",
    "email": "john@doe.com",
    "password": "password"
}'
```

## Best Practices

- **Security:**
  JWT authentication ensures that only authorized users can access and modify their to-do items.
- **Error Handling:**
  The API provides meaningful HTTP status codes and error responses for validation issues and unauthorized access.

- **Modular Architecture:**
  With a clear separation between controllers, services, repositories, and middlewares, the application is maintainable and testable.

- **Swagger Integration:**
  Auto-generated Swagger documentation facilitates easy API exploration and integration testing.

## License

This project is open source and available under the [MIT License](LICENSE).

Happy coding!
