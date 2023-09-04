package _struct

import (
	GrpcUserService "server/MainService/GrpcClients/UserService"
)

type RegisterUserResponse struct {
	UserId int32 `json:"UserId"`
}

func GetRegisterUserResponse(resp *GrpcUserService.RegisterUserResponse) *RegisterUserResponse {
	return &RegisterUserResponse{
		UserId: resp.UserId,
	}
}

type RegisterTutorResponse struct {
	TutorId int32 `json:"TutorId"`
}
