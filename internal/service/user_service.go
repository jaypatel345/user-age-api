package service

import (
	"context"
	"time"

	"github.com/jaypatel/user-age-api/internal/models"

	db "github.com/jaypatel/user-age-api/db/sqlc"
	"github.com/jaypatel/user-age-api/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(r *repository.UserRepository) *UserService {
	return &UserService{
		repo: r,
	}
}

func (s *UserService) CreateUser(
	ctx context.Context,
	params db.CreateUserParams,
) (db.User, error) {
	return s.repo.CreateUser(ctx, params)
}

func (s *UserService) UpdateUser(
	ctx context.Context,
	params db.UpdateUserParams,
) (db.User, error) {
	return s.repo.UpdateUser(ctx, params)
}

func (s *UserService) DeleteUser(
	ctx context.Context,
	id int32,
) error {
	return s.repo.DeleteUser(ctx, id)
}

func (s *UserService) GetUserByID(
	ctx context.Context,
	id int32,
) (models.UserResponse, error) {
	user, err := s.repo.GetUserByID(ctx, id)
	if err != nil {
		return models.UserResponse{}, err
	}

	return models.UserResponse{
		ID:   user.ID,
		Name: user.Name,
		DOB:  user.Dob.Format("2006-01-02"),
		Age:  calculateAge(user.Dob),
	}, nil
}

func (s *UserService) ListUsers(ctx context.Context,
) ([]models.UserResponse, error) {
	users, err := s.repo.ListUsers(ctx)

	if err != nil {
		return nil, err
	}

	response := make([]models.UserResponse, 0, len(users))

	for _, user := range users {
		response = append(response, models.UserResponse{
			ID:   user.ID,
			Name: user.Name,
			DOB:  user.Dob.Format("2006-01-02"),
			Age:  calculateAge(user.Dob),
		})
	}
	return response, nil
}

func calculateAge(dob time.Time) int {
	today := time.Now()

	age := today.Year() - dob.Year()

	if today.Month() < dob.Month() ||

		(today.Month() == dob.Month() &&
			today.Day() < dob.Day()) {
		age--
	}
	return age
}
