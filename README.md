# Connect - Go Web Application

## Overview
Connect is a Go-based web application that provides user authentication with signup and login functionality, and a protected dashboard. It uses JWT for session management and MySQL as the database.

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

## License
This project is licensed under the MIT License.
