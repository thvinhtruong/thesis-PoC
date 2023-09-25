package GrpcStudyService

import context "context"

type ZStudyServiceClient struct {
	innerClient StudyServiceClient
}

func NewZStudyServiceClient(innerClient StudyServiceClient) *ZStudyServiceClient {
	return &ZStudyServiceClient{innerClient}
}

func (c *ZStudyServiceClient) GetUserRecord(request *GetUserRecordRequest) *GetUserRecordResponse {
	response, _ := c.innerClient.GetUserRecord(context.Background(), request)
	return response
}
