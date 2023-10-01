package GrpcStudyService

import (
	context "context"
	"server/MainService/config"
)

type ZStudyServiceClient struct {
	innerClient StudyServiceClient
	Config      config.Config
}

func (c *ZStudyServiceClient) CreateUserRecord(request *CreateUserRecordRequest) *CreateUserRecordResponse {
	response, _ := c.innerClient.CreateUserRecord(context.Background(), request)
	return response
}

func (c *ZStudyServiceClient) GetUserRecord(request *GetUserRecordRequest) *GetUserRecordResponse {
	response, _ := c.innerClient.GetUserRecord(context.Background(), request)
	return response
}
