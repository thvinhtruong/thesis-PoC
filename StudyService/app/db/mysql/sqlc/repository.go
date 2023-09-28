package db

import (
	"context"
	"database/sql"
)

type StudyRepository interface {
	GetUserRecordTx(ctx context.Context, arg GetUserRecordParams) (GetUserRecordResult, error)
}

func NewStudyRepository(db *sql.DB) StudyRepository {
	return &TxStore{
		Queries: New(db),
		db:      db,
	}
}
