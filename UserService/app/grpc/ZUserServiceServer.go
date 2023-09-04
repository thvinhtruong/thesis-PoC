package GrpcUserService

import (
	"context"
	"server/UserService/app/apperror"
	db "server/UserService/app/db/mysql/sqlc"
)

type ZUserServiceServer struct {
	repository db.Repository
}

func NewZUserServiceServer(repository db.Repository) ZUserServiceServer {
	return ZUserServiceServer{repository: repository}
}

func (s *ZUserServiceServer) LoginUser(ctx context.Context, request *LoginUserRequest) (*LoginUserResponse, error) {
	user := db.LoginUserParams{
		Email:    request.Phone,
		Password: request.Password,
	}

	record, err := s.repository.LoginUser(ctx, user)
	if err != nil {
		return &LoginUserResponse{
			UserId:    -1,
			ErrorCode: apperror.GetCode(err),
		}, err
	}

	return &LoginUserResponse{
		UserId:    record.ID,
		ErrorCode: apperror.GetCode(err),
	}, nil
}

func (s *ZUserServiceServer) RegisterUser(ctx context.Context, request *RegisterUserRequest) (*RegisterUserResponse, error) {
	user := db.RegisterUserParams{
		Fullname: request.Fullname,
		Password: request.Password,
		Email:    request.Phone,
		Gender:   request.Gender,
	}

	record, err := s.repository.RegisterUser(ctx, user)
	if err != nil {
		return &RegisterUserResponse{
			UserId:    -1,
			ErrorCode: apperror.GetCode(err),
		}, err
	}

	return &RegisterUserResponse{
		UserId:    record.ID,
		ErrorCode: apperror.GetCode(err),
	}, nil
}

func (s *ZUserServiceServer) mustEmbedUnimplementedUserServiceServer() {
}
