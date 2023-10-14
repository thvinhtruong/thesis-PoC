package GrpcStudyService

import (
	context "context"
	"log"
	"server/StudyService/config"
)

type ZStudyServiceClient struct {
	innerClient StudyServiceClient
	Config      config.Config
}

func NewZStudyServiceClient(innerClient StudyServiceClient) *ZStudyServiceClient {
	return &ZStudyServiceClient{innerClient, nil}
}

func (c *ZStudyServiceClient) CreateUserRecord(request *CreateUserRecordRequest) *CreateUserRecordResponse {
	response, err := c.innerClient.CreateUserRecord(context.Background(), request)
	if err != nil {
		log.Printf("Error creating user record: %v", err)
		return nil
	}
	return response
}

func (c *ZStudyServiceClient) GetUserRecord(request *GetUserRecordRequest) *GetUserRecordResponse {
	response, err := c.innerClient.GetUserRecord(context.Background(), request)
	if err != nil {
		log.Printf("Error getting user record: %v", err)
		return nil
	}
	return response
}
