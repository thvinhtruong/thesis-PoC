package GrpcStudyService

import context "context"

type ZStudyServiceClient struct {
	innerClient StudyServiceClient
}

func NewZStudyServiceClient(innerClient StudyServiceClient) *ZStudyServiceClient {
	return &ZStudyServiceClient{innerClient}
}

func (c *ZStudyServiceClient) CreateUserRecord(request *CreateUserRecordRequest) *CreateUserRecordResponse {
	response, _ := c.innerClient.CreateUserRecord(context.Background(), request)
	return response
}

func (c *ZStudyServiceClient) GetUserRecord(request *GetUserRecordRequest) *GetUserRecordResponse {
	response, _ := c.innerClient.GetUserRecord(context.Background(), request)
	return response
}
