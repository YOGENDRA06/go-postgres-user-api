# User Management API (Go Backend Development Task)

A RESTful API built using **Go** to manage users with `name` and `date of birth (dob)`.
The API dynamically calculates and returns a user’s **age** when fetching user details.

---


##  Setup & Run (Docker – Recommended)

### 1. Clone Repository
```bash
git clone <your-repo-url>
cd Go_Backend_Development_Task
```

### 2. Build & Run
```bash
docker compose up --build
```

This will:
- Start PostgreSQL
- Apply migrations
- Start the API server

---

##  Run Locally (Without Docker)

```bash
go mod tidy
sqlc generate
go run cmd/server/main.go
```

---

##  API Endpoints

### Create User
`POST /api/users`

```json
{
  "name": "John Doe",
  "dob": "1998-05-10"
}
```

### Get All Users
`GET /api/users`

### Get User by ID
`GET /api/users/{id}`

---

##  Age Calculation

Age is calculated dynamically using the current date and `dob`.
It is **not stored** in the database.

---

##  Notes

- Ensure database migrations run successfully.

- If you see `relation "users" does not exist`, re-check migrations.

---

##  Author

Yogendra Kumar