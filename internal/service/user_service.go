package service

import (
	"context"
	"time"

	"ainyx/db/sqlc"
	"ainyx/internal/models"
	"ainyx/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(
	repo *repository.UserRepository,
) *UserService {

	return &UserService{
		repo: repo,
	}
}

func (s *UserService) CreateUser(
	ctx context.Context,
	name string,
	dob string,
) (sqlc.User, error) {

	parsedDOB, err := time.Parse(
		"2006-01-02",
		dob,
	)

	if err != nil {
		return sqlc.User{}, err
	}

	return s.repo.CreateUser(
		ctx,
		sqlc.CreateUserParams{
			Name: name,
			Dob:  parsedDOB,
		},
	)
}

func (s *UserService) GetUser(
	ctx context.Context,
	id int32,
) (models.UserResponse, error) {

	user, err := s.repo.GetUser(ctx, id)
	if err != nil {
		return models.UserResponse{}, err
	}

	return models.UserResponse{
		ID:   user.ID,
		Name: user.Name,
		DOB:  user.Dob.Format("2006-01-02"),
		Age:  CalculateAge(user.Dob),
	}, nil
}

func (s *UserService) GetAllUsers(
	ctx context.Context,
) ([]models.UserResponse, error) {

	users, err := s.repo.ListUsers(ctx)
	if err != nil {
		return nil, err
	}

	var response []models.UserResponse

	for _, user := range users {
		response = append(
			response,
			models.UserResponse{
				ID:   user.ID,
				Name: user.Name,
				DOB:  user.Dob.Format("2006-01-02"),
				Age:  CalculateAge(user.Dob),
			},
		)
	}

	return response, nil
}

func (s *UserService) UpdateUser(
	ctx context.Context,
	id int32,
	name string,
	dob string,
) (sqlc.User, error) {

	parsedDOB, err := time.Parse(
		"2006-01-02",
		dob,
	)

	if err != nil {
		return sqlc.User{}, err
	}

	return s.repo.UpdateUser(
		ctx,
		sqlc.UpdateUserParams{
			ID:   id,
			Name: name,
			Dob:  parsedDOB,
		},
	)
}

func (s *UserService) DeleteUser(
	ctx context.Context,
	id int32,
) error {

	return s.repo.DeleteUser(
		ctx,
		id,
	)
}
