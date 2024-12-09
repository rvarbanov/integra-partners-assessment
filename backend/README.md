# User Management API

Backend service for user management system built with Go 1.22 and Echo framework.

## Prerequisites

- Go 1.22 or higher
- PostgreSQL
- Make

## API Documentation

### Endpoints

- `GET /api/v1/status` - Get server status
- `GET /api/v1/users` - List all users
- `POST /api/v1/user` - Create a new user
- `GET /api/v1/user/:id` - Get a user by id
- `PUT /api/v1/user/:id` - Update an existing user
- `DELETE /api/v1/user/:id` - Delete a user

Detailed API documentation is available at `/api/v1/docs/index.html` when the server is running.

### Example Requests

#### Get Status
Request:
```bash
curl -X GET http://localhost:8080/api/v1/status
```

Response:
```json
{
  "status": "ok"
}
```

#### Create User
Request:
```bash
curl -X POST http://localhost:8080/api/v1/user \
-H "Content-Type: application/json" \
-d '{
  "user_name": "testuser",
  "first_name": "John",
  "last_name": "Doe",
  "email": "testuser@example.com",
  "user_status": "A",
  "department": "Engineering"
}'
```

Response:
```json
{
  "data": {
    "user_ id": 1, 
    "user_name": "testuser",
    "first_name": "John",
    "last_name": "Doe",
    "email": "testuser@example.com",
    "user_status": "A",
    "department": "Engineering"
  }
}
```

#### Update User
Request:
```bash
curl -X PUT http://localhost:8080/api/v1/user/:id \
-H "Content-Type: application/json" \
-d '{
  "user_name": "testuser",
  "first_name": "John",
  "last_name": "Doe",
  "email": "testuser@example.com",
  "user_status": "A",
  "department": "Engineering"
}'
```

Response:
```json
{
  "data": {
    "user_id": 1, 
    "user_name": "testuser",
    "first_name": "John",
    "last_name": "Doe",
    "email": "testuser@example.com",
    "user_status": "A",
    "department": "Engineering"
  }
}
```

#### Delete User
Request:
```bash
curl -X DELETE http://localhost:8080/api/v1/user/:id
```

#### Get User
Request:
```bash
curl -X GET http://localhost:8080/api/v1/user/:id
```

Response:
```json
{
  "data": {
    "user_id": 1, 
    "user_name": "testuser",
    "first_name": "John",
    "last_name": "Doe",
    "email": "testuser@example.com",
    "user_status": "A",
    "department": "Engineering"
  }
}
```

#### List Users
Request:
```bash
curl -X GET http://localhost:8080/api/v1/users
```
Response:
```json
{
  "data": [
    {
      "user_id": 1, 
      "user_name": "testuser",
      "first_name": "John",
      "last_name": "Doe",
      "email": "testuser@example.com",
      "user_status": "A",
      "department": "Engineering"
    },
    {
      ...
    },
    {
      ...
    }
  ]
}
```

## Error Handling

The API uses standard HTTP status codes:

- 200: Success
- 400: Bad Request (invalid input)
- 404: Not Found
- 409: Conflict (e.g., username already exists)
- 500: Internal Server Error

Error responses follow this format:

```json
{
  "error": "Error message here"
}
```

## Development Guidelines

1. Follow [Uber Go Style Guide](https://github.com/uber-go/guide/blob/master/style.md)
2. Write tests for new functionality
3. Update swagger documentation for API changes

## Dependencies

- [Echo](https://echo.labstack.com/) - Web framework
- [Squirrel](https://github.com/Masterminds/squirrel) - SQL query builder
- [Ginkgo](https://github.com/onsi/ginkgo) - Testing framework
- [Gomega](https://github.com/onsi/gomega) - Matcher/assertion library
- [Swag](https://github.com/swaggo/swag) - Swagger documentation generator

## Build and Run

```bash
make build
make run
```

## DB Migration
After the app is build and run, you can run the migration to create the user table.

```bash
make migrate
```

## API Swagger Documentation
```bash
make doc
```

## Testing

Run tests with coverage report:

```bash
make mock
make test
```

## Improvement Opportunities

- Add logging
  - Log errors, warnings, info, debug
- Add metrics
  - Integrate with Datadog
- Add authentication
  - Implement JWT-based authentication
  - Update API endpoints to require authentication
