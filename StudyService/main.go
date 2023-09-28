package main

import (
	"fmt"
	"log"
	"net"
	db "server/StudyService/app/db/mysql/sqlc"
	GrpcStudyService "server/StudyService/app/grpc"
	config "server/StudyService/config"
	"server/UserService/app/sqlconnection"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	configuration := config.Singleton
	dbconn := sqlconnection.DBConn
	repository := db.NewStudyRepository(dbconn)
	host := fmt.Sprintf(
		"%v:%v",
		configuration.GetConfig(config.STUDY_SERVICE_HOST),
		configuration.GetConfig(config.STUDY_SERVICE_PORT))

	lis, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatalf("Failed to listen: %v", host)
	}

	s := grpc.NewServer()
	studyServiceServer := GrpcStudyService.NewZStudyServiceServer(repository)

	GrpcStudyService.RegisterStudyServiceServer(s, &studyServiceServer)
	reflection.Register(s)

	fmt.Printf("StudyService is listening at %v", host)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to start server %v", err)
	}
}
