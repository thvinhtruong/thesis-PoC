package GrpcStudyService

import (
	context "context"
	"fmt"
	"log"
	"server/MainService/config"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var StudyInstance *ZStudyServiceClient

func init() {
	Config := config.GetInstance()
	target := fmt.Sprintf(
		"%v:%v",
		Config.GetConfig(config.STUDY_SERVICE_HOST),
		Config.GetConfig(config.STUDY_SERVICE_PORT),
	)
	log.Println("Connecting ... to StudyService at", target)
	conn, err := grpc.Dial(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Connect to StudyService failed")
		log.Fatal(err)
	}

	innerClient := NewStudyServiceClient(conn)
	StudyInstance = &ZStudyServiceClient{
		innerClient: innerClient,
		Config:      Config,
	}

	fmt.Println("Connected to StudyService success")
}

type ZStudyServiceClient struct {
	innerClient StudyServiceClient
	Config      config.Config
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
