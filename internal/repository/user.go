package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	pw "authservice/internal/password"
)

type IUserObject interface {
	GetUserByUsername(username string) (UserObject, error)
	CreateUser(username string, password string) error
}

type UserObjectRepository struct {
	connection *pgx.Conn
}

type UserObject struct {
	UserId   int64
	Username string
	PassHash string
}

func (repository *UserObjectRepository) GetUserByUsername(username string) (UserObject, error) {
	sqlRequest := fmt.Sprintf("select id, password_hash from users where username = '%s';", username)
	row := repository.connection.QueryRow(context.Background(), sqlRequest)
	var id int64
	var passwordHash string
	row.Scan(&id, &passwordHash)
	user := UserObject{id, username, passwordHash}
	if len(user.PassHash) > 0{
		return user, nil
	}
	return UserObject{}, fmt.Errorf("user %s does not exist", user.Username)
}

func (repository *UserObjectRepository) CreateUser(username string, password string) error {
	passHash, hashErr := pw.HashPassword(password)
	if hashErr != nil{
		return fmt.Errorf("password didn't hashed")
	}
	sqlRequest := fmt.Sprintf("insert into users (username, password_hash) values ('%s', '%s')", username, passHash)
	_, execErr := repository.connection.Exec(context.Background(), sqlRequest)
	if execErr != nil{
		return fmt.Errorf("error with inserting in table: %s", execErr)
	}
	return nil
}

func NewUserObjectRepository(connection *pgx.Conn) IUserObject {
	return &UserObjectRepository{connection}
}
