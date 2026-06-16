# user-age-api

A Go REST API for storing users and calculating their age from date of birth. The project uses Fiber, PostgreSQL, sqlc, and a handler/service/repository structure.

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

## Common Errors

`DATABASE_URL is required`

The `.env` file is missing or does not contain `DATABASE_URL`.

`pq: role "username" does not exist`

The database URL uses a PostgreSQL username that does not exist locally.

`sql: no rows in result set`

The requested user ID does not exist. Create a user first or use an ID returned by `GET /users`.
