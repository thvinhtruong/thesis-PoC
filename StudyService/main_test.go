package main

import (
	"context"
	"os"
	db "server/StudyService/app/db/mysql/sqlc"
	GrpcStudyService "server/StudyService/app/grpc"
	"server/StudyService/app/sqlconnection"
	"testing"
)

func TestMain(m *testing.M) {
	dbconn := sqlconnection.GetDB()
	if dbconn == nil {
		panic("database is nil")
	}
	if err := dbconn.Ping(); err != nil {
		panic(err)
	}

	os.Exit(m.Run())
}

func TestGetUserRecord(t *testing.T) {
	dbConn := sqlconnection.GetDB()

	repository := db.NewStudyRepository(dbConn)

	studyServiceServer := GrpcStudyService.NewZStudyServiceServer(repository)
	request := GrpcStudyService.GetUserRecordRequest{
		UserId: 1,
	}

	result, err := studyServiceServer.GetUserRecord(context.Background(), &request)

	if err != nil {
		t.Fatal(err)
	}

	if result == nil {
		t.Fatal("result is nil")
	}

	if len(result.UserRecord) == 0 {
		t.Fatal("result is empty")
	}

	for _, item := range result.UserRecord {
		if item.Name == "" {
			t.Fatal("name is empty")
		}
		if item.Weight == 0 {
			t.Fatal("weight is empty")
		}
		if item.Score == 0 {
			t.Fatal("score is empty")
		}
	}

	// log.Println(result)
}
