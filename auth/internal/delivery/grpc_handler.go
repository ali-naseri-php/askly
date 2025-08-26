package delivery

import (
	"auth/internal/service"

	authpb "github.com/ali-naseri-php/Askly/proto/auth"
	"context"
)

type AuthHandler struct {
	authpb.UnimplementedAuthServiceServer
	svc *service.AuthService
}

func NewAuthHandler(svc *service.AuthService) *AuthHandler {
	return &AuthHandler{svc: svc}
}

func (h *AuthHandler) Register(ctx context.Context, req *authpb.RegisterRequest) (*authpb.RegisterResponse, error) {
	user, err := h.svc.Register(ctx, req.Email, req.Password)
	if err != nil {
		return &authpb.RegisterResponse{
			UserId:  "",
			Message: err.Error(),
		}, nil
	}
	return &authpb.RegisterResponse{
		UserId:  user.ID,
		Message: "registration successful",
	}, nil
}

func (h *AuthHandler) Login(ctx context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	token, err := h.svc.Login(ctx, req.Email, req.Password)
	if err != nil {
		return &authpb.LoginResponse{
			Token:   "",
			Message: err.Error(),
		}, nil
	}
	return &authpb.LoginResponse{
		Token:   token,
		Message: "login successful",
	}, nil
}
