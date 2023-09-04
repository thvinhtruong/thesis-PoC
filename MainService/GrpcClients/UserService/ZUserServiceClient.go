package GrpcUserService

import (
	"context"
	"fmt"
	"log"
	"server/MainService/config"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var Instance *ZUserServiceClient

func init() {
	Config := config.GetInstance()
	target := fmt.Sprintf(
		"%v:%v",
		Config.GetConfig(config.USER_SERVICE_HOST),
		Config.GetConfig(config.USER_SERVICE_PORT),
	)
	log.Println("Connecting ... to UserService at", target)
	conn, err := grpc.Dial(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Connect to UserService failed")
		log.Fatal(err)
	}

	innerClient := NewUserServiceClient(conn)
	Instance = &ZUserServiceClient{
		innerClient: innerClient,
		Config:      Config,
	}

	fmt.Println("Connected to UserService success")
}

type ZUserServiceClient struct {
	innerClient UserServiceClient
	Config      config.Config
}

func (c *ZUserServiceClient) LoginUser(request *LoginUserRequest) *LoginUserResponse {
	response, _ := c.innerClient.LoginUser(context.Background(), request)
	return response
}

func (c *ZUserServiceClient) RegisterUser(request *RegisterUserRequest) *RegisterUserResponse {
	response, _ := c.innerClient.RegisterUser(context.Background(), request)
	return response
}
