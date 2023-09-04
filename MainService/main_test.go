package main

import (
	"fmt"
	"log"
	GrpcUserService "server/MainService/GrpcClients/UserService"
	"server/MainService/config"
	"server/MainService/handlers"
	"testing"

	"github.com/stretchr/testify/require"
)

func setupSuite(tb testing.TB) func(tb testing.TB) {
	return func(tb testing.TB) {
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

	got := userHandler.Repo.RegisterUser(&GrpcUserService.RegisterUserRequest{
		Fullname: "Bao",
		Phone:    "0934777111",
		Password: "Tam",
		Gender:   "Male",
	})

	log.Printf("got: %v", got)

	require.NotNil(t, got)
}
