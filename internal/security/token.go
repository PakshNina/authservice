package security

import (
	"authservice/internal/config"
	password2 "authservice/internal/password"
	rp "authservice/internal/repository"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func CreateToken(conf config.Config, password string, user rp.UserObject) (string, error) {
	canGenerateToken := password2.CheckPasswordHash(password, user.PassHash)
	if canGenerateToken == false {
		return "", fmt.Errorf("wrong password for user %s", user.Username)
	}
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = user.UserId
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(conf.AccessSecret))
	if err != nil {
		return "", err
	}
	return token, nil
}