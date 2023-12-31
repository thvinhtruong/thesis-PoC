package main

import (
	"context"
	"log"
	"os"
	db "server/UserService/app/db/mysql/sqlc"
	GrpcUserService "server/UserService/app/grpc"
	"server/UserService/app/sqlconnection"
	"server/UserService/pkg/random"
	"testing"

	"github.com/stretchr/testify/require"
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

func TestRegisterUser(t *testing.T) {
	dbconn := sqlconnection.GetDB()
	repository := db.NewRepository(dbconn)

	user := db.RegisterUserParams{
		Email:    random.RandomPhone(),
		Password: "123456789",
		Fullname: "Vinh",
		Gender:   random.RandomGender(),
	}

	userServiceServer := GrpcUserService.NewZUserServiceServer(repository)
	request := GrpcUserService.RegisterUserRequest{
		Email:    user.Email,
		Password: user.Password,
		Fullname: user.Fullname,
		Gender:   user.Gender,
	}

	result, err := userServiceServer.RegisterUser(context.Background(), &request)
	log.Println(result)
	require.NoError(t, err)
	require.NotEmpty(t, result)
}
