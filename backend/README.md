# User Management API

Backend service for user management system built with Go 1.22 and Echo framework.

## Prerequisites

- Go 1.22 or higher
- PostgreSQL
- Make

## API Documentation

### Endpoints

- `GET /status` - Get server status
- `GET /users` - List all users
- `POST /user` - Create a new user
- `PUT /user/:id` - Update an existing user
- `DELETE /user/:id` - Delete a user

Detailed API documentation is available at `/swagger/index.html` when the server is running.

### Example Requests

#### Get Status
Request:
```bash
curl -X GET http://localhost:8080/status
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
curl -X POST http://localhost:8080/user \
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
    "id": 1, 
    "user_name": "testuser",
    "first_name": "John",
    "last_name": "Doe",
    "email": "testuser@example.com",
    "user_status": "A",
    "department": "Engineering"
  }
}
``

#### Update User
Request:
```bash
curl -X PUT http://localhost:8080/user/:id \
-H "Content-Type: application/json" \
-d '{
  "user_name": "testuser",
  "first_name": "John",
  "last_name": "Doe",
  "email": "testuser@example.com",
  "user_status": "A",
  "department": "Engineering",
}'
```
Response:
```json
{
  "data": {
    "id": 1, 
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
curl -X DELETE http://localhost:8080/user/:id
```

#### Get User
Request:
```bash
curl -X GET http://localhost:8080/user/:id
```
Response:
```json
{
  "data": {
    "id": 1, 
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
curl -X GET http://localhost:8080/users
```
Response:
```json
{
  "data": [
    {
      "id": 1, 
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

## Testing

Run tests with coverage report:

```bash
make mock
make test
```

##TODO:
- Add db migration
  - create db user/password
  - create the user table
- Add docker-compose
  - testing

## Improvement opportunities:
- Add logging
  - log errors, warnings, info, debug
- Add metrics
  - datadog
- Add authentication
