package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	GrpcUserService "server/MainService/GrpcClients/UserService"
	"server/MainService/config"
	"server/MainService/handlers"
)

func main() {
	configuration := config.GetInstance()
	userHandler := handlers.NewUserApiHanlder(configuration, GrpcUserService.Instance)
	http.HandleFunc("/api/LoginUser", userHandler.LoginUser)
	http.HandleFunc("/api/RegisterUser", userHandler.RegisterUser)
	dir, _ := os.Getwd()
	fmt.Println("current path :" + dir)

	port := configuration.GetConfig(config.MAIN_SERVICE_PORT)

	fmt.Printf(
		"Server is listening at %v:%v\n",
		configuration.GetConfig(config.MAIN_SERVICE_HOST),
		configuration.GetConfig(config.MAIN_SERVICE_PORT),
	)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
