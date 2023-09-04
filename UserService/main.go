package main

import (
	"fmt"
	"log"
	"net"
	db "server/UserService/app/db/mysql/sqlc"
	GrpcUserService "server/UserService/app/grpc"
	"server/UserService/app/sqlconnection"
	config "server/UserService/config"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	configuration := config.Singleton
	dbconn := sqlconnection.DBConn
	repository := db.NewRepository(dbconn)
	host := fmt.Sprintf(
		"%v:%v",
		configuration.GetConfig(config.USER_SERVICE_HOST),
		configuration.GetConfig(config.USER_SERVICE_PORT))

	lis, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatalf("Failed to listen: %v", host)
	}

	s := grpc.NewServer()
	userServiceServer := GrpcUserService.NewZUserServiceServer(repository)

	GrpcUserService.RegisterUserServiceServer(s, &userServiceServer)
	reflection.Register(s)

	fmt.Printf("ZUserServiceServer is listening at %v", host)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to start server %v", err)
	}
}
