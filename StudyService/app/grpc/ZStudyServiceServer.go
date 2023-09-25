package GrpcStudyService

import (
	context "context"
	db "server/StudyService/app/db/mysql/sqlc"
)

type ZStudyServiceServer struct {
	innerServer db.StudyRepository
}

func NewZStudyServiceServer(innerServer db.StudyRepository) ZStudyServiceServer {
	return ZStudyServiceServer{innerServer: innerServer}
}

func (s *ZStudyServiceServer) GetUserRecord(ctx context.Context, request *GetUserRecordRequest) (*GetUserRecordResponse, error) {
	var response *GetUserRecordResponse
	resquest := db.GetUserRecordParams{
		UserId: request.UserId,
	}

	record, err := s.innerServer.GetUserRecord(ctx, resquest)
	if err != nil {
		return nil, err
	}

	for pos, item := range response.UserRecord {
		item.Name = record.UserRecord[pos].Name
		item.Weight = record.UserRecord[pos].Weight
		item.Score = record.UserRecord[pos].Score
	}

	return response, nil
}

func (s *ZStudyServiceServer) mustEmbedUnimplementedStudyServiceServer() {
}
