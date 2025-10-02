# Connect - Go Web Application Template

## Overview
My_Web_Template_Of_Connect is a Go-based web application template that provides user authentication with signup and login functionality, and a protected dashboard. It uses JWT for session management and MySQL as the database. This repository is designed to be used as a starting point for new Go web projects.

## Features
- User signup with validation and password hashing
- User login with JWT authentication and secure cookies
- Protected dashboard accessible only to authenticated users
- HTML templates for login, signup, and dashboard pages
- Environment variable configuration and MySQL database integration

## Technologies Used
- Go 1.24
- MySQL
- JWT for authentication
- bcrypt for password hashing
- HTML templates for frontend rendering

## Getting Started

### Prerequisites
- Go 1.24 or higher installed
- MySQL database setup
- Environment variables configured (see `.env`)

### Installation
1. Clone the repository
2. Set up your `.env` file with database credentials and JWT secret
3. Run `go mod download` to install dependencies
4. Start the application:
   ```
   go run main.go
   ```
5. Access the app at `http://localhost:8080`

## Project Structure
- `main.go`: Application entry point, route registration, and server start
- `handler/`: HTTP handlers for signup, login, middleware, and dashboard
- `config/`: Environment loading and database initialization
- `db/`: Database interaction functions
- `templates/`: HTML templates for rendering pages
- `.gitignore`: Specifies files and folders to ignore in Git

## Using as a Template

This repository is set up as a GitHub template. To use it for a new project:

1. On GitHub, click the "Use this template" button to create a new repository.
2. Clone the new repository to your local machine.
3. Update the following files for your project:
   - `go.mod`: Change the module name to match your new project.
   - `.env`: Configure your database credentials and JWT secret.
   - `templates/`: Customize HTML templates as needed.
   - `handler/`: Modify or add handlers for your specific requirements.
4. Run `go mod tidy` to update dependencies.
5. Start developing your application.

This template provides a solid foundation with authentication, database integration, and a basic dashboard, allowing you to focus on building your specific features.

## License
This project is licensed under the MIT License.
