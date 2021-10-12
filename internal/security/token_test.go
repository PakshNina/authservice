package security

import (
	"authservice/internal/config"
	rp "authservice/internal/repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateToken(t *testing.T) {
	cfg := config.Config{AccessSecret: "Secret"}
	password := "P@ssword"
	user := rp.UserObject{UserId: 1, Username: "user", PassHash: "$2a$14$I8EjWVFx9k5zyogucf4b7ePRofKmG0ioqb0hbvrU/AgkntVk8v7v6"}
	token, err := CreateToken(cfg, password, user)
	assert.NotNil(t, token)
	assert.Nil(t, err)
}
