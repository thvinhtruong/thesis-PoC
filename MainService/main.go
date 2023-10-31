package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	GrpcStudyService "server/MainService/GrpcClients/StudyService"
	GrpcUserService "server/MainService/GrpcClients/UserService"
	"server/MainService/config"
	"server/MainService/handlers"
	"server/MainService/middleware"

	"github.com/gorilla/mux"
)

func main() {
	// router
	router := mux.NewRouter()

	// config
	configuration := config.GetInstance()

	// handlers
	userHandler := handlers.NewUserApiHanlder(configuration, GrpcUserService.Instance)
	studyHandler := handlers.NewStudyHandler(configuration, GrpcStudyService.StudyInstance)

	// // middleware
	// if configuration.GetConfig(config.ENABLE_HTTP_CACHE) == "true" {
	// 	log.Println("Enable http cache")
	// 	router.Use(reverseproxy.HttpResponseCachingMiddleware)
	// }

	// routing
	router.HandleFunc("/api/v1/LoginUser", userHandler.LoginUser)
	router.HandleFunc("/api/v1/RegisterUser", userHandler.RegisterUser)
	router.HandleFunc("/api/v1/GetUserRecord/{id}", studyHandler.GetUserRecord)
	router.HandleFunc("/api/v1/CreateUserRecord", studyHandler.CreateUserRecord)

	// Calculate post request
	router.Use(middleware.MeasureResponseDuration)

	dir, _ := os.Getwd()
	fmt.Println("current path :" + dir)

	port := configuration.GetConfig(config.MAIN_SERVICE_PORT)

	fmt.Printf(
		"Server is listening at %v:%v\n",
		configuration.GetConfig(config.MAIN_SERVICE_HOST),
		configuration.GetConfig(config.MAIN_SERVICE_PORT),
	)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), router))
}
