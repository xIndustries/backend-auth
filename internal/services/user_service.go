package services

import (
	"context"
	"errors"
	"log"
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

// Helper functions
func stringPtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func derefString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

// ✅ CreateUser - Prevent duplicate creation
func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserResponse, error) {
	log.Printf("🔹 Checking if user exists | Auth0ID: %s", req.Auth0Id)

	existingUser, err := s.Repo.GetUser(req.Auth0Id)
	if err == nil && existingUser != nil {
		log.Printf("✅ User already exists, skipping creation | Auth0ID: %s", req.Auth0Id)
		return &pb.UserResponse{
			Id:        existingUser.ID,
			Auth0Id:   existingUser.Auth0ID,
			Email:     existingUser.Email,
			Username:  derefString(existingUser.Username),
			CreatedAt: existingUser.CreatedAt.Format(time.RFC3339),
		}, nil
	}

	log.Printf("🔹 Creating new user | Auth0ID: %s | Email: %s", req.Auth0Id, req.Email)

	user := &models.User{
		ID:        uuid.NewString(),
		Auth0ID:   req.Auth0Id,
		Email:     req.Email,
		Username:  stringPtr(req.Username),
		CreatedAt: time.Now(),
	}

	err = s.Repo.CreateUser(user)
	if err != nil {
		log.Printf("❌ Failed to create user: %v", err)
		return nil, err
	}

	log.Printf("✅ User created successfully: %s", user.ID)

	return &pb.UserResponse{
		Id:        user.ID,
		Auth0Id:   user.Auth0ID,
		Email:     user.Email,
		Username:  derefString(user.Username),
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
	}, nil
}

// ✅ GetUser
func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.UserResponse, error) {
	log.Printf("🔹 Retrieving user | Auth0ID: %s", req.Auth0Id)

	user, err := s.Repo.GetUser(req.Auth0Id)
	if err != nil {
		log.Printf("❌ Failed to retrieve user: %v", err)
		return nil, err
	}

	log.Printf("✅ User retrieved successfully: %s", user.Auth0ID)

	return &pb.UserResponse{
		Id:        user.ID,
		Auth0Id:   user.Auth0ID,
		Email:     user.Email,
		Username:  derefString(user.Username),
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
	}, nil
}

// ✅ UpdateUsername
func (s *UserService) UpdateUsername(ctx context.Context, req *pb.UpdateUsernameRequest) (*pb.UserResponse, error) {
	log.Printf("🔹 Updating username | Auth0ID: %s | New Username: %s", req.Auth0Id, req.Username)

	if req.Username == "" {
		log.Println("❌ UpdateUsername: Username is empty")
		return nil, errors.New("username is required")
	}

	err := s.Repo.UpdateUsername(req.Auth0Id, req.Username)
	if err != nil {
		log.Printf("❌ Failed to update username in DB: %v", err)
		return nil, err
	}

	log.Println("✅ Username updated successfully in DB")

	user, err := s.Repo.GetUser(req.Auth0Id)
	if err != nil {
		log.Printf("❌ Failed to retrieve updated user: %v", err)
		return nil, err
	}

	log.Printf("✅ Username update confirmed | Auth0ID: %s | Username: %s", user.Auth0ID, derefString(user.Username))

	return &pb.UserResponse{
		Id:        user.ID,
		Auth0Id:   user.Auth0ID,
		Email:     user.Email,
		Username:  derefString(user.Username),
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
	}, nil
}

// ✅ UpdateUser (Email)
func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UserResponse, error) {
	log.Printf("🔹 Updating user email | Auth0ID: %s | New Email: %s", req.Auth0Id, req.Email)

	if req.Email == "" {
		log.Println("❌ UpdateUser: Email is empty")
		return nil, errors.New("email is required")
	}

	err := s.Repo.UpdateUserEmail(req.Auth0Id, req.Email)
	if err != nil {
		log.Printf("❌ Failed to update email: %v", err)
		return nil, err
	}

	log.Println("✅ Email updated successfully")

	user, err := s.Repo.GetUser(req.Auth0Id)
	if err != nil {
		log.Printf("❌ Failed to retrieve updated user: %v", err)
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

// ✅ DeleteUser
func (s *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	log.Printf("🔹 Deleting user | Auth0ID: %s", req.Auth0Id)

	err := s.Repo.DeleteUser(req.Auth0Id)
	if err != nil {
		log.Printf("❌ Failed to delete user: %v", err)
		return nil, err
	}

	log.Printf("✅ User deleted successfully: %s", req.Auth0Id)

	return &pb.DeleteUserResponse{
		Message: "User deleted successfully",
	}, nil
}
