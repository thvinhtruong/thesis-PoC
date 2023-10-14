package main

import (
	"fmt"
	"log"
	GrpcStudyService "server/MainService/GrpcClients/StudyService"
	GrpcUserService "server/MainService/GrpcClients/UserService"
	"server/MainService/config"
	"server/MainService/handlers"
	"server/UserService/pkg/random"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func setupSuite(tb testing.TB) func(tb testing.TB) {
	return func(tb testing.TB) {
	}
}

func TDD_RandomName() string {
	return "Student " + strconv.Itoa(random.RandomInt(1, 1000))
}

func TDD_CreateRandomUser() GrpcUserService.RegisterUserRequest {
	return GrpcUserService.RegisterUserRequest{
		Fullname: TDD_RandomName(),
		Email:    random.RandomEmail(),
		Password: random.RandomString(10),
		Gender:   random.RandomGender(),
	}
}

func TestRegisterUser(t *testing.T) {
	t.Logf("TestRegisterUser\n")
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)

	configuration := config.GetInstance()
	userHandler := handlers.NewUserApiHanlder(configuration, GrpcUserService.Instance)

	fmt.Printf(
		"Server is listening at %v:%v\n",
		configuration.GetConfig(config.MAIN_SERVICE_HOST),
		configuration.GetConfig(config.MAIN_SERVICE_PORT),
	)

	var createUserRequest = TDD_CreateRandomUser()
	response := userHandler.Repo.RegisterUser(&createUserRequest)

	log.Printf("response: %v", response)

	require.NotNil(t, response)
}

func TestGetUserRecord(t *testing.T) {
	t.Logf("TestGetUserRecord\n")
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)

	configuration := config.GetInstance()
	studyHandler := handlers.NewStudyHandler(configuration, GrpcStudyService.StudyInstance)

	fmt.Printf(
		"Server is listening at %v:%v\n",
		configuration.GetConfig(config.MAIN_SERVICE_HOST),
		configuration.GetConfig(config.MAIN_SERVICE_PORT),
	)

	got := studyHandler.Repo.GetUserRecord(&GrpcStudyService.GetUserRecordRequest{
		UserId: 1,
	})

	log.Printf("got: %v", got)

	require.NotNil(t, got)
}
