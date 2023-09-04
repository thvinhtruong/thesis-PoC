package GrpcUserService

import (
	"context"
)

type ZUserServiceClient struct {
	innerClient UserServiceClient
}

func (c *ZUserServiceClient) LoginUser(request *LoginUserRequest) *LoginUserResponse {
	response, _ := c.innerClient.LoginUser(context.Background(), request)
	return response
}

func (c *ZUserServiceClient) RegisterUser(request *RegisterUserRequest) *RegisterUserResponse {
	response, _ := c.innerClient.RegisterUser(context.Background(), request)
	return response
}
