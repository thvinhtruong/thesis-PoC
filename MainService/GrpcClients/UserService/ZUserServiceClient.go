package GrpcUserService

import (
	"context"
	"log"
	"server/MainService/config"
)

type ZUserServiceClient struct {
	innerClient UserServiceClient
	Config      *config.SConfig
}

func (c *ZUserServiceClient) LoginUser(request *LoginUserRequest) *LoginUserResponse {
	response, err := c.innerClient.LoginUser(context.Background(), request)
	if err != nil {
		log.Printf("Error logging in user: %v", err)
		return nil
	}
	return response
}

func (c *ZUserServiceClient) RegisterUser(request *RegisterUserRequest) *RegisterUserResponse {
	response, err := c.innerClient.RegisterUser(context.Background(), request)
	if err != nil {
		log.Printf("Error registering user: %v", err)
		return nil
	}
	return response
}
