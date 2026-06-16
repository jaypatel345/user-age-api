# user-age-api

A small Go API project for storing users and their date of birth in PostgreSQL.

## What is implemented so far

- Created a PostgreSQL `users` table schema in `db/migrations/001_create_users.sql`.
- Added sqlc queries in `db/queries/users.sql`.
- Configured `sqlc.yaml` to generate Go database code.
- Generated sqlc code into `db/sqlc`.
- Added a repository layer in `internal/repository/user_repository.go`.

## Users table

The `users` table has:

```sql
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    dob DATE NOT NULL
);
```

## sqlc

sqlc reads:

- schema from `db/migrations`
- queries from `db/queries`
- generated Go code output to `db/sqlc`

To regenerate database code after changing SQL files:

```bash
sqlc generate
```

## Repository

`internal/repository/user_repository.go` wraps the generated sqlc methods:

- `CreateUser`
- `UpdateUser`
- `GetUserByID`
- `ListUsers`
- `DeleteUser`

The repository uses:

```go
db "github.com/jaypatel/user-age-api/db/sqlc"
```

`dob` is generated as `time.Time` because the database column type is `DATE`.
