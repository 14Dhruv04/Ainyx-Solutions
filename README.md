# User Management REST API

This is my submission for the task assigned by Ainyx Solutions for the recruitment of SDE Intern role.

The objective was to build a RESTful API in Go that manages users and their DOB while dynamically calculating and returning a user's age whenever user information is fetched.

## Tech Stack
- Go + Fiber for backend
- PostgreSQL + SQLC for database
- Uber Zap for logging
- go-playground/validator for validation
- Go Testing Package for testing

## Architecture
This project follows a layered architecture where an HTTP request is initially parsed by the **Handler** Layer which then calls the **Service** layer that contains the main logic to execute tasks. The service layer then requests the **Repository** layer which is responsible for database operations and delegation of queries to SQLC generated code.

### Database Schema
```
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    dob DATE NOT NULL
);
```
Only the user's Date of Birth is stored, ensuring Age is never stored rather calculated dynamically whenever user data is fetched.

### Age Calculation Algorithm
- Calculate the difference between the current year and birth year.
- Check whether the user's birthday has occurred in the current year.
- Adjust hte result if the birthday has not yet occurred.

## Timeline

I was initially unfamiliar with Go, Fiber, SQLC and several tools mentioned in the tech stack and hence resorted to the documentation of each tool and ChatGPT. 
- I started with establishing the project structure that was already mentioned in the given task and then installing the tools required for implementation purposes.
- Structured the database schema and wrote SQL querires for CRUD operations.
- Generated the database access layer using SQLC.
- Implemented the repository layer and then the business logic in the service layer.
- Implemented the REST handler and routes using Fiber.
- Added dynamic age calculation.
- Added request validation.
- Added structured logging.
- Added request duration logging middleware.
- Added unit tests for the age calculation function.
- Performed testing using PowerShell requests and PostgreSQL verification.

## How to Run

Clone Repository
```
git clone <https://github.com/14Dhruv04/Ainyx-Solutions.git>
cd Ainyx-Solutions
```

Update the database connection string
```
$env:DB_URL=postgres://postgres:<YOUR_PASSWORD>@localhost:5432/userdb?sslmode=disable
```

Generate SQLC Code
```
sqlc generate
```

Run Application
```
go run cmd/server/main.go
```

The Server starts on ```http://localhost:3000```

## Learnings
- Go and the similarity it shares with C++ and Java.
- Building REST APIs using Fiber
- Dependency Injection
- Request Validation
- Middleware Implementation
