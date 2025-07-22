# Workout Plan API

A RESTful API built with Go for managing workout plans and exercises. This project uses Prisma Client Go for database access and Gorilla Mux for routing.

## Features
- Create, read, update, and delete workout plans
- Each workout plan can have multiple exercises
- Built with Go, Prisma, and PostgreSQL

## Requirements
- Go 1.24.4 or later
- PostgreSQL database

## Getting Started

### 1. Clone the repository
```bash
git clone https://github.com/sibelephant/workout-plan-api.git
cd workout-plan-api
```

### 2. Set up environment variables
Create a `.env` file in the project root with the following:
```
DATABASE_URL=postgresql://USER:PASSWORD@HOST:PORT/DATABASE
```
Replace `USER`, `PASSWORD`, `HOST`, `PORT`, and `DATABASE` with your PostgreSQL credentials.

### 3. Install dependencies
```bash
go mod download
```

### 4. Run database migrations
(Assuming you have Prisma CLI installed and your schema is up to date)
```bash
prisma migrate deploy
```

### 5. Run the API server
```bash
go run ./cmd/api/main.go
```
The server will start on `http://localhost:8080`.

## API Endpoints

### Workout Plans
- `POST   /workout-plans` — Create a new workout plan
- `GET    /workout-plans` — Get all workout plans
- `GET    /workout-plans/{id}` — Get a specific workout plan (currently mapped to all plans, may need adjustment)
- `PUT    /workout-plans/{id}` — Update a workout plan by ID
- `DELETE /workout-plans/{id}` — Delete a workout plan by ID

#### WorkoutPlan JSON structure
```json
{
  "id": "string",
  "name": "string",
  "description": "string",
  "exercises": [
    {
      "id": "string",
      "name": "string",
      "sets": 0,
      "reps": 0,
      "workout_plan_id": "string"
    }
  ]
}
```

## Project Structure
- `cmd/api/main.go` — Entry point for the API server
- `internal/handlers/` — HTTP handler functions
- `internal/models/` — Data models
- `internal/database/` — Database connection logic
- `prisma/` — Prisma schema and generated client

## Development
- Use Go modules for dependency management
- Prisma Client Go is used for database access
- Environment variables are loaded from `.env`

## License
MIT
