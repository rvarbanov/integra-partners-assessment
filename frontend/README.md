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

## Running end-to-end tests

Run `ng e2e` to execute the end-to-end tests via a platform of your choice. To use this command, you need to first add a package that implements end-to-end testing capabilities.

## Further help

To get more help on the Angular CLI use `ng help` or go check out the [Angular CLI Overview and Command Reference](https://angular.io/cli) page.

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

- `GET /users` - List all users
- `POST /user` - Create new user
- `PUT /user/:id` - Update user
- `DELETE /user/:id` - Delete user
- `GET /user/:id` - Get single user

## TODO

- [ ] Add error handling service
- [ ] Implement loading states
- [ ] Add unit tests
- [ ] Add E2E tests with Cypress
- [ ] Add proper error messages
- [ ] Add proper loading states
- [ ] Add proper success messages
- [ ] Add proper confirmation dialogs
- [ ] Add proper form validation
- [ ] Add proper routing
- [ ] Add proper documentation