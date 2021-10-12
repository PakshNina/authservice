package server

import (
	"authservice/internal/config"
	"authservice/internal/repository"
	"authservice/internal/security"
	login "authservice/pkg/pb/proto"
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type service struct {
	repository repository.IUserObject
	conf       config.Config
	login.UnimplementedLoginServiceServer
}

func NewService(repository repository.IUserObject, conf config.Config) *service {
	return &service{
		repository: repository,
		conf: conf,
		UnimplementedLoginServiceServer: login.UnimplementedLoginServiceServer{},
	}
}

func (s *service) Login(ctx context.Context, req *login.LoginRequest) (*login.LoginResponse, error) {
	user, hashErr := s.repository.GetUserByUsername("user")
	if hashErr != nil{
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("username %s was not found", req.Username),
		)
	}
	accessToken, tokenErr := security.CreateToken(s.conf, req.Password, user)
	if tokenErr != nil{
		return nil, status.Errorf(codes.PermissionDenied, fmt.Sprintf("%s", tokenErr))
	}
	response := new(login.LoginResponse)
	response.AccessToken = accessToken
	return response, nil
}
