package services

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/xIndustries/BandRoom/backend-auth/internal/models"
	"github.com/xIndustries/BandRoom/backend-auth/internal/repositories"
	pb "github.com/xIndustries/BandRoom/backend-auth/proto/Generated"
)

type UserService struct {
	Repo *repositories.UserRepository
}

// NewUserService creates a new UserService instance.
func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

// Helper function to convert string to *string
func stringPtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

// Helper function to convert *string to string
func derefString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

// CreateUser handles the creation of a new user.
func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserResponse, error) {
	user := &models.User{
		ID:        uuid.NewString(),
		Auth0ID:   req.Auth0Id,
		Email:     req.Email,
		Username:  stringPtr(req.Username), // Convert string to *string
		CreatedAt: time.Now(),
	}

	err := s.Repo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return &pb.UserResponse{
		Id:        user.ID,
		Auth0Id:   user.Auth0ID,
		Email:     user.Email,
		Username:  derefString(user.Username), // Convert *string to string
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
	}, nil
}

// GetUser handles retrieving a user by Auth0 ID.
func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.UserResponse, error) {
	user, err := s.Repo.GetUser(req.Auth0Id)
	if err != nil {
		return nil, err
	}

	return &pb.UserResponse{
		Id:        user.ID,
		Auth0Id:   user.Auth0ID,
		Email:     user.Email,
		Username:  derefString(user.Username), // Convert *string to string
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
	}, nil
}

// UpdateUser handles updating a user's email.
func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UserResponse, error) {
	if req.Email == "" {
		return nil, errors.New("email is required")
	}

	err := s.Repo.UpdateUserEmail(req.Auth0Id, req.Email)
	if err != nil {
		return nil, err
	}

	user, err := s.Repo.GetUser(req.Auth0Id)
	if err != nil {
		return nil, err
	}

	return &pb.UserResponse{
		Id:        user.ID,
		Auth0Id:   user.Auth0ID,
		Email:     user.Email,
		Username:  derefString(user.Username),
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
	}, nil
}

// UpdateUsername handles updating a user's username.
func (s *UserService) UpdateUsername(ctx context.Context, req *pb.UpdateUsernameRequest) (*pb.UserResponse, error) {
	if req.Username == "" {
		return nil, errors.New("username is required")
	}

	err := s.Repo.UpdateUsername(req.Auth0Id, req.Username)
	if err != nil {
		return nil, err
	}

	user, err := s.Repo.GetUser(req.Auth0Id)
	if err != nil {
		return nil, err
	}

	return &pb.UserResponse{
		Id:        user.ID,
		Auth0Id:   user.Auth0ID,
		Email:     user.Email,
		Username:  derefString(user.Username),
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
	}, nil
}

// DeleteUser handles deleting a user by Auth0 ID.
func (s *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	err := s.Repo.DeleteUser(req.Auth0Id)
	if err != nil {
		return nil, err
	}

	return &pb.DeleteUserResponse{
		Message: "User deleted successfully",
	}, nil
}
