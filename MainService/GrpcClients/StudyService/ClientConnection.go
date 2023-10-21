package GrpcStudyService

import (
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
