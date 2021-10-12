package server

import (
	"authservice/internal/config"
	"authservice/internal/repository"
	login "authservice/pkg/pb/proto"
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc/status"
	"testing"
)

type mockRepository struct {
	mock.Mock
}

func (mock *mockRepository) GetUserByUsername(username string) (repository.UserObject, error) {
	args := mock.Called(username)
	return args.Get(0).(repository.UserObject), args.Error(1)
}
func (mock *mockRepository) CreateUser(username string, password string) error {
	args := mock.Called(username, password)
	return args.Error(0)
}

func TestServiceLogin(t *testing.T) {
	mockRepository := new(mockRepository)
	conf := config.Config{}
	hashPassword := "$2a$14$I8EjWVFx9k5zyogucf4b7ePRofKmG0ioqb0hbvrU/AgkntVk8v7v6"
	user := repository.UserObject{UserId: 1, Username: "user", PassHash: hashPassword}
	pbRequest := &login.LoginRequest{Username: "user", Password: "P@ssword"}
	mockRepository.On("GetUserByUsername", "user").Return(user, nil)
	service := NewService(mockRepository, conf)
	accessToken, err := service.Login(context.Background(), pbRequest)
	errStatus := status.Convert(err)
	assert.Nil(t, errStatus)
	assert.NotEmpty(t, accessToken.AccessToken)
}