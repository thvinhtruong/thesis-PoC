package _struct

import (
	GrpcStudyService "server/MainService/GrpcClients/StudyService"
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

type UserRecordRow struct {
	Subject string `json:"Subject"`
	Weight  int32  `json:"Weight"`
	Score   int32  `json:"Score"`
}

type UserRecordResponse struct {
	UserId     int             `json:"UserId"`
	UserRecord []UserRecordRow `json:"UserRecord"`
}

func GetUserRecordResponseData(userID int, resp *GrpcStudyService.GetUserRecordResponse) UserRecordResponse {
	userStudyRecordRows := make([]UserRecordRow, len(resp.UserRecord))

	for i, row := range resp.UserRecord {
		userStudyRecordRows[i] = UserRecordRow{
			Subject: row.Name,
			Weight:  row.Weight,
			Score:   row.Score,
		}
	}

	return UserRecordResponse{
		UserId:     userID,
		UserRecord: userStudyRecordRows,
	}
}
