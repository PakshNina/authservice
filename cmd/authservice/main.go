package main

import (
	"authservice/internal/config"
	"authservice/internal/db"
	"authservice/internal/repository"
	login "authservice/internal/server"
	proto "authservice/pkg/pb/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	cfg, confErr := config.NewConfig()
	if confErr != nil{
		log.Fatalf("failed to config: %v", confErr)
	}
	listener, listenErr := net.Listen("tcp", cfg.ServerAddress)
	if listenErr != nil {
		log.Fatalf("failed to listen: %v", listenErr)
	}
	connection, connErr := db.InitPostgreConnection(*cfg)
	if connErr != nil {
		log.Fatalf("failed to connect to db: %v", connErr)
	}
	userRepository := repository.NewUserObjectRepository(connection)

	authServer := grpc.NewServer()
	service := login.NewService(userRepository, *cfg)
	proto.RegisterLoginServiceServer(authServer, service)
	log.Printf("server listening at %v", listener.Addr())
	if serverErr := authServer.Serve(listener); serverErr != nil {
		log.Fatalf("failed to serve: %v", serverErr)
	}
}
