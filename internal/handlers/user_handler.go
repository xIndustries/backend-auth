package handlers

import (
	"context"

	"github.com/xIndustries/BandRoom/backend-auth/internal/services"
	pb "github.com/xIndustries/BandRoom/backend-auth/proto/Generated"
)

type UserHandler struct {
	Service *services.UserService
	pb.UnimplementedUserServiceServer
}

// NewUserHandler creates a new UserHandler instance.
func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{Service: service}
}

func (h *UserHandler) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserResponse, error) {
	return h.Service.CreateUser(ctx, req)
}

func (h *UserHandler) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.UserResponse, error) {
	return h.Service.GetUser(ctx, req)
}

func (h *UserHandler) UpdateUsername(ctx context.Context, req *pb.UpdateUsernameRequest) (*pb.UserResponse, error) {
	return h.Service.UpdateUsername(ctx, req)
}

func (h *UserHandler) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UserResponse, error) {
	return h.Service.UpdateUser(ctx, req)
}

func (h *UserHandler) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	return h.Service.DeleteUser(ctx, req)
}
