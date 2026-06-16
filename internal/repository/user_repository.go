package repository

import (
	"context"

	db "github.com/jaypatel/user-age-api/db/sqlc"
)

type UserRepository struct {
	queries *db.Queries
}

func NewUserRepository(q *db.Queries) *UserRepository {
	return &UserRepository{
		queries: q,
	}
}

func (r *UserRepository) CreateUser(
	ctx context.Context,
	params db.CreateUserParams,
) (db.User, error) {

	return r.queries.CreateUser(ctx, params)
}

func (r *UserRepository) UpdateUser(
	ctx context.Context,
	params db.UpdateUserParams,
) (db.User, error) {

	return r.queries.UpdateUser(ctx, params)
}

func (r *UserRepository) GetUserByID(
	ctx context.Context,
	id int32,
) (db.User, error) {
	return r.queries.GetUser(ctx, id)
}

func (r *UserRepository) ListUsers(ctx context.Context,
) ([]db.User, error) {
	return r.queries.ListUsers(ctx)
}

func (r *UserRepository) DeleteUser(
	ctx context.Context,
	id int32,
) error {
	return r.queries.DeleteUser(ctx, id)

}
