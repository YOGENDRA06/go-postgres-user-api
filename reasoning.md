# Reasoning & Design Decisions

## Overview

This project was built as part of the Go Backend Development Task for the Software Engineering Intern role.  
The goal was to design and implement a clean, maintainable RESTful API in Go that manages users with their name and date of birth, and dynamically calculates age when fetching user data.

The solution strictly follows the technology stack and guidelines mentioned in the task description.

---

## Tech Stack Selection

The following stack was used as required:

- **Go (Golang)** – Primary programming language
- **GoFiber** – Lightweight and high-performance HTTP framework
- **PostgreSQL** – Relational database for persistence
- **SQLC** – Type-safe SQL query generation
- **Zap Logger** – Structured and performant logging
- **go-playground/validator** – Input validation
- **Docker & Docker Compose** – Environment consistency and easy setup

Each tool was chosen to keep the solution production-ready while remaining simple and easy to reason about.

---

## Project Architecture

The project follows a layered architecture to ensure separation of concerns:

- **Handler Layer** – Handles HTTP requests and responses
- **Service Layer** – Contains business logic (e.g., age calculation)
- **Repository Layer** – Handles database operations using SQLC
- **Routes Layer** – Centralized route definitions
- **Middleware Layer** – Request logging, request ID handling, etc.
- **Models** – Shared data structures
- **Logger** – Centralized logging configuration

This structure improves readability, testability, and long-term maintainability.

---

## Database Design

A single `users` table is used with the following core fields:
- `id` (UUID)
- `name`
- `dob`

The **age is not stored** in the database.  
Instead, it is calculated dynamically at runtime to avoid redundant data and ensure accuracy over time.

On application startup, the service checks whether the required table exists and creates it if missing.  
This ensures smooth local and containerized execution without requiring manual database setup.

---

## Age Calculation Logic

Age is calculated in the service layer using the current date and the user's date of birth.

This approach:
- Keeps the database schema simple
- Ensures the age is always accurate
- Avoids data inconsistency

---

## Validation & Error Handling

- Incoming requests are validated using `go-playground/validator`
- Proper HTTP status codes are returned for validation and server errors
- Errors are logged using Zap with structured fields for better observability

---

## Logging Strategy

Zap logger is used to:
- Log incoming requests
- Capture errors with stack traces
- Provide consistent and structured logs

This helps in debugging and production monitoring.

---

## SQLC Usage

SQLC is used to generate type-safe database access code from raw SQL queries.

Benefits:
- Compile-time query validation
- No ORM overhead
- Clear separation between SQL and application logic

---

## Dockerization

Docker and Docker Compose are used to:
- Run PostgreSQL and the API together
- Ensure consistent environments across machines
- Simplify setup and execution

The application can be started with a single command:
```bash
docker compose up --build
```

---

## Key Design Decisions

- **Layered architecture** for clarity and scalability
- **Dynamic age calculation** instead of storing derived data
- **Type-safe DB access** using SQLC
- **Startup DB initialization** to reduce manual setup
- **Structured logging** for observability

---

## Conclusion

The solution focuses on clarity, correctness, and maintainability while adhering strictly to the task requirements.  
The design choices were made to reflect real-world backend development practices in Go, with an emphasis on clean architecture and explainable decisions.
