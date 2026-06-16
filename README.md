# user-age-api

A small Go REST API for storing users with their date of birth in PostgreSQL. The API uses Fiber for HTTP routing, sqlc for generated database queries, and a simple handler/service/repository structure.

## Features

- Create a user with name and date of birth
- List all users with calculated age
- Get one user by ID with calculated age
- Update a user
- Delete a user
- Load database configuration from `.env`

## Tech Stack

- Go
- Fiber
- PostgreSQL
- sqlc
- lib/pq PostgreSQL driver
- godotenv

## Project Structure

```text
cmd/server/main.go              # App entrypoint
db/migrations/                  # Database migrations
db/queries/                     # sqlc SQL queries
db/sqlc/                        # Generated sqlc code
internal/handler/               # HTTP handlers
internal/routes/                # Fiber routes
internal/service/               # Business logic
internal/repository/            # Database access wrapper
internal/models/                # Response models
```

## Database Setup

Create the PostgreSQL database:

```bash
createdb user_age_db
```

Run the migration:

```bash
psql user_age_db -f db/migrations/001_create_users.sql
```

The migration creates this table:

```sql
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    dob DATE NOT NULL
);
```

## Environment Setup

Create a `.env` file in the project root:

```env
DATABASE_URL=postgres://localhost:5432/user_age_db?sslmode=disable
```

If your PostgreSQL setup needs a username and password:

```env
DATABASE_URL=postgres://username:password@localhost:5432/user_age_db?sslmode=disable
```

`.env` is ignored by Git, so local secrets should not be pushed to GitHub.

## Run The Server

```bash
go run ./cmd/server
```

The server starts on:

```text
http://localhost:8080
```

## API Routes

```text
POST    /users
GET     /users
GET     /users/:id
PUT     /users/:id
DELETE  /users/:id
```

## Test With Postman

Use this header for requests with JSON bodies:

```text
Content-Type: application/json
```

### Create User

```http
POST http://localhost:8080/users
```

Body:

```json
{
  "name": "Jay Patel",
  "dob": "2000-05-20"
}
```

Example response:

```json
{
  "id": 1,
  "name": "Jay Patel",
  "dob": "2000-05-20T00:00:00Z"
}
```

### List Users

```http
GET http://localhost:8080/users
```

Example response:

```json
[
  {
    "id": 1,
    "name": "Jay Patel",
    "dob": "2000-05-20",
    "age": 26
  }
]
```

### Get User By ID

```http
GET http://localhost:8080/users/1
```

Example response:

```json
{
  "id": 1,
  "name": "Jay Patel",
  "dob": "2000-05-20",
  "age": 26
}
```

### Update User

```http
PUT http://localhost:8080/users/1
```

Body:

```json
{
  "name": "Jay Patel Updated",
  "dob": "1999-12-10"
}
```

Example response:

```json
{
  "id": 1,
  "name": "Jay Patel Updated",
  "dob": "1999-12-10T00:00:00Z"
}
```

### Delete User

```http
DELETE http://localhost:8080/users/1
```

Expected response:

```text
204 No Content
```

## Regenerate sqlc Code

If you change SQL files in `db/queries` or migrations used by sqlc, regenerate code with:

```bash
sqlc generate
```

## Common Errors

`DATABASE_URL is required`

The `.env` file is missing or does not contain `DATABASE_URL`.

`pq: role "username" does not exist`

Your database URL uses a PostgreSQL username that does not exist locally. Use your local PostgreSQL role or use the no-password local URL if your setup allows it.

`sql: no rows in result set`

The requested user ID does not exist. Create a user first or use an ID returned by `GET /users`.
