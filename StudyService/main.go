package main

import (
	"fmt"
	"log"
	"net"
	db "server/StudyService/app/db/mysql/sqlc"
	GrpcStudyService "server/StudyService/app/grpc"
	"server/StudyService/app/sqlconnection"
	config "server/StudyService/config"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	configuration := config.Singleton
	dbconn := sqlconnection.GetDB()
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
