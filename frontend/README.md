# User Management UI

This project was generated with [Angular CLI](https://github.com/angular/angular-cli) version 16.2.16.

## Development server

Run `ng serve` for a dev server. Navigate to `http://localhost:4200/`. The application will automatically reload if you change any of the source files.

## Code scaffolding

Run `ng generate component component-name` to generate a new component. You can also use `ng generate directive|pipe|service|class|guard|interface|enum|module`.

## Build

Run `ng build` to build the project. The build artifacts will be stored in the `dist/` directory.

## Running unit tests

Run `ng test` to execute the unit tests via [Karma](https://karma-runner.github.io).

## Features

- User Management Interface
  - List all users in a data grid
  - Create new users
  - Update existing users
  - Delete users
- Form validation
  - Required fields validation
  - Email format validation
  - Display API error messages
- Routing for all views/dialogs
- Follows Angular best practices and style guide

## API Integration

The frontend communicates with the backend API running on `http://localhost:8080`. API endpoints:

- `GET /api/v1/users` - List all users
- `POST /api/v1/user` - Create new user
- `PUT /api/v1/user/:id` - Update user
- `DELETE /api/v1/user/:id` - Delete user
- `GET /api/v1/user/:id` - Get single user
