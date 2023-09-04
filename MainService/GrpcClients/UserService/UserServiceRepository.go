package GrpcUserService

type UserServiceRepository interface {
	RegisterUser(req *RegisterUserRequest) *RegisterUserResponse
	LoginUser(req *LoginUserRequest) *LoginUserResponse
}
