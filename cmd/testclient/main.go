package main

import (
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
