package main

import (
	"authservice/internal/config"
	"authservice/internal/db"
	"authservice/internal/repository"
	proto "authservice/pkg/pb/proto"
	"context"
	"flag"
	"google.golang.org/grpc"
	"log"
)

func main() {
	username := flag.String("username", "user", "username")
	password := flag.String("password", "P@ssword", "password for user")

	flag.Parse()
	cfg, confErr := config.NewConfig()
	if confErr != nil{
		log.Fatalf("failed to config: %v", confErr)
	}
	dbConn, dbConnErr := db.InitPostgreConnection(*cfg)
	if dbConnErr != nil {
		log.Fatalf("failed to connect to db: %v", dbConnErr)
	}
	userRepository := repository.NewUserObjectRepository(dbConn)
	// Create user to make sure it will work
	userRepository.CreateUser(*username, *password)
	connection, connErr := grpc.Dial("127.0.0.1:8888", grpc.WithInsecure())
	if connErr != nil {
		log.Fatalf("Error: %v", connErr)
	}
	client := proto.NewLoginServiceClient(connection)
	resp, loginErr := client.Login(context.Background(),
		&proto.LoginRequest{Username: *username, Password: *password})
	if loginErr != nil {
		log.Fatalf("Error: %v", loginErr)
	}
	log.Println("Access token:", resp.AccessToken)
}
