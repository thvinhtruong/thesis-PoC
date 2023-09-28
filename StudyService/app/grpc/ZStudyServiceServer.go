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
	var userStudyRecordRows []UserStudyRecord
	resquest := db.GetUserRecordParams{
		UserId: request.UserId,
	}

	record, err := s.innerServer.GetUserRecordTx(ctx, resquest)
	if err != nil {
		return nil, err
	}

	recordRow := record.UserRecord

	for pos, item := range recordRow {
		userStudyRecordRows[pos].Name = item.Name
		userStudyRecordRows[pos].Weight = item.Weight
		userStudyRecordRows[pos].Score = item.Score
	}

	return response, nil
}

func (s *ZStudyServiceServer) mustEmbedUnimplementedStudyServiceServer() {
}
