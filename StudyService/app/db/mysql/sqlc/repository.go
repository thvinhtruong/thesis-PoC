package db

import (
	"context"
	"database/sql"
)

type StudyRepository interface {
	CreateUserRecordTx(ctx context.Context, arg CreateUserRecordParams) (CreateUserRecordResult, error)
	GetUserRecordTx(ctx context.Context, arg GetUserRecordParams) (GetUserRecordResult, error)
}

func NewStudyRepository(db *sql.DB) StudyRepository {
	return &TxStore{
		Queries: New(db),
		db:      db,
	}
}
