package GrpcStudyService

type StudyServiceRepository interface {
	CreateUserRecord(request *CreateUserRecordRequest) *CreateUserRecordResponse
	GetUserRecord(request *GetUserRecordRequest) *GetUserRecordResponse
}
