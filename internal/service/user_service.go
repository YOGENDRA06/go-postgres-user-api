package service

import (
	"context"
	"errors"
	"time"

	"Go_Backend_Development_Task/db/sqlc"
	"Go_Backend_Development_Task/internal/models"
	"Go_Backend_Development_Task/internal/repository"
)

var ErrInvalidDOB = errors.New("invalid dob format")

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// ---------------- CREATE USER ----------------
func (s *UserService) CreateUser(
	ctx context.Context,
	req models.CreateUserRequest,
) (*models.UserResponse, error) {

	dob, err := time.Parse("2006-01-02", req.DOB)
	if err != nil {
		return nil, ErrInvalidDOB
	}

	user, err := s.repo.CreateUser(ctx, req.Name, dob)
	if err != nil {
		return nil, err
	}

	return s.buildResponse(user), nil
}

// ---------------- GET USER BY ID ----------------
func (s *UserService) GetUser(
	ctx context.Context,
	id int64,
) (*models.UserResponse, error) {

	user, err := s.repo.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return s.buildResponse(user), nil
}

// ---------------- GET USERS (PAGINATED) ----------------
func (s *UserService) GetUsers(
	ctx context.Context,
	page int,
	limit int,
) ([]models.UserResponse, error) {

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 50 {
		limit = 10
	}

	offset := (page - 1) * limit

	users, err := s.repo.GetUsers(
		ctx,
		int32(limit),
		int32(offset),
	)
	if err != nil {
		return nil, err
	}

	res := make([]models.UserResponse, 0, len(users))
	for _, u := range users {
		res = append(res, *s.buildResponse(u))
	}

	return res, nil
}

// ---------------- UPDATE USER ----------------
func (s *UserService) UpdateUser(
	ctx context.Context,
	id int64,
	req models.CreateUserRequest,
) (*models.UserResponse, error) {

	dob, err := time.Parse("2006-01-02", req.DOB)
	if err != nil {
		return nil, ErrInvalidDOB
	}

	user, err := s.repo.UpdateUser(ctx, id, req.Name, dob)
	if err != nil {
		return nil, err
	}

	return s.buildResponse(user), nil
}

// ---------------- DELETE USER ----------------
func (s *UserService) DeleteUser(
	ctx context.Context,
	id int64,
) error {
	return s.repo.DeleteUser(ctx, id)
}

// ---------------- HELPERS ----------------
func (s *UserService) buildResponse(
	u sqlc.User,
) *models.UserResponse {

	return &models.UserResponse{
		ID:   int64(u.ID),
		Name: u.Name,
		DOB:  u.Dob.Format("2006-01-02"),
		Age:  calculateAge(u.Dob),
	}
}

func calculateAge(dob time.Time) int {
	now := time.Now()
	age := now.Year() - dob.Year()
	if now.YearDay() < dob.YearDay() {
		age--
	}
	return age
}
