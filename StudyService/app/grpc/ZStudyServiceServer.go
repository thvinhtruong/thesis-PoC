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
	resquest := db.GetUserRecordParams{
		UserId: request.UserId,
	}

	record, err := s.innerServer.GetUserRecordTx(ctx, resquest)
	if err != nil {
		return nil, err
	}

	recordRow := record.UserRecord

	userStudyRecordRows := make([]*UserStudyRecord, len(recordRow))

	for i, row := range recordRow {
		userStudyRecordRows[i] = &UserStudyRecord{
			Name:   row.Name,
			Weight: row.Weight,
			Score:  row.Score,
		}
	}

	response := &GetUserRecordResponse{
		UserRecord: userStudyRecordRows,
	}

	return response, nil
}

func (s *ZStudyServiceServer) CreateUserRecord(ctx context.Context, request *CreateUserRecordRequest) (*CreateUserRecordResponse, error) {
	requestParams := db.CreateUserRecordParams{
		UserID:   request.UserId,
		ModuleID: request.ModuleId,
		Weight:   request.Weight,
		Score:    request.Score,
	}

	record, err := s.innerServer.CreateUserRecordTx(ctx, requestParams)
	if err != nil {
		return nil, err
	}

	response := &CreateUserRecordResponse{
		UserId: record.UserId,
	}

	return response, nil
}

func (s *ZStudyServiceServer) mustEmbedUnimplementedStudyServiceServer() {
}
