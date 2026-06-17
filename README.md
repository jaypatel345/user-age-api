# user-age-api

A Go REST API for storing users and calculating their age from date of birth. The project uses Fiber, PostgreSQL, sqlc, Zap logging, request middleware, and a handler/service/repository structure.

## Final Checklist

- Clean folder structure
- sqlc database queries working
- CRUD APIs complete
- Age calculated dynamically from DOB
- Request validation added
- Zap logger integrated
- Request ID and request logging middleware added
- Proper HTTP status codes
- README written

## Setup Steps

1. Clone the repository and go into the project folder:

```bash
git clone https://github.com/jaypatel345/user-age-api.git
cd user-age-api
```

2. Install Go dependencies:

```bash
go mod tidy
```

3. Create the PostgreSQL database:

```bash
createdb user_age_db
```

4. Run the database migration:

```bash
psql user_age_db -f db/migrations/001_create_users.sql
```

5. Create a `.env` file from the example:

```bash
cp .env.example .env
```

6. Update `.env` if your local PostgreSQL connection is different.

## How To Run

Start the API server:

```bash
go run ./cmd/server
```

The server runs at:

```text
http://localhost:8080
```

Run tests:

```bash
go test ./...
```

Regenerate sqlc code after SQL changes:

```bash
sqlc generate
```

## Project Structure

```text
cmd/server/main.go              # App entrypoint
db/migrations/                  # Database migrations
db/queries/                     # sqlc SQL queries
db/sqlc/                        # Generated sqlc code
internal/handler/               # HTTP handlers
internal/logger/                # Zap logger setup
internal/middleware/            # Request ID and logging middleware
internal/models/                # Response models
internal/repository/            # Database access wrapper
internal/routes/                # Fiber routes
internal/service/               # Business logic and age calculation
```

## Env Variables

Create a `.env` file in the project root.

```env
DATABASE_URL=postgres://localhost:5432/user_age_db?sslmode=disable
```

If your PostgreSQL setup uses username and password:

```env
DATABASE_URL=postgres://username:password@localhost:5432/user_age_db?sslmode=disable
```

`.env` is ignored by Git. Keep real local database credentials there, and commit only `.env.example`.

## API Docs

Base URL:

```text
http://localhost:8080
```

For requests with JSON bodies, use:

```text
Content-Type: application/json
```

Every response includes:

```text
X-Request-ID: <uuid>
```

Validation rules:

- `name` is required for create and update.
- `dob` must use `YYYY-MM-DD` format.
- `id` path params must be valid integers.

Status codes:

- `200 OK` for successful reads and updates
- `201 Created` for successful create
- `204 No Content` for successful delete
- `400 Bad Request` for invalid request body, ID, name, or DOB
- `404 Not Found` when a user ID does not exist
- `500 Internal Server Error` for unexpected server/database errors

### Create User

```http
POST /users
```

Request body:

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

Validation error example:

```json
{
  "error": "name is required"
}
```

### List Users

```http
GET /users
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
GET /users/1
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
PUT /users/1
```

Request body:

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
DELETE /users/1
```

Expected response:

```text
204 No Content
```

### Not Found Example

```http
GET /users/999
```

Example response:

```json
{
  "error": "user not found"
}
```

## Common Errors

`DATABASE_URL is required`

The `.env` file is missing or does not contain `DATABASE_URL`.

`pq: role "username" does not exist`

The database URL uses a PostgreSQL username that does not exist locally.

`sql: no rows in result set`

The requested user ID does not exist. The API now returns `404 user not found` for this case.
