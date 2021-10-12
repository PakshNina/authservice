package db

import (
	"authservice/internal/config"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
)

func InitPostgreConnection(conf config.Config) (*pgx.Conn, error) {
	var dbURL = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", conf.User, conf.Password, conf.DbHost, conf.DbPort, conf.DbName)
	connection, connErr := pgx.Connect(context.Background(), dbURL)
	return connection, connErr
}
