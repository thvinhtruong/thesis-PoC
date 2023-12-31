// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package db

import (
	"context"
	"database/sql"
)

type Querier interface {
	CreateModule(ctx context.Context, arg CreateModuleParams) (sql.Result, error)
	CreateUserRecord(ctx context.Context, arg CreateUserRecordParams) error
	DeleteModule(ctx context.Context, id int32) error
	DeleteUserRecord(ctx context.Context, id int32) error
	GetModule(ctx context.Context, id int32) (Module, error)
	GetUserRecord(ctx context.Context, userID int32) ([]GetUserRecordRow, error)
	GetUserRecordById(ctx context.Context, id int32) (Userrecord, error)
	GetUserRecordByUserIdAndModuleId(ctx context.Context, arg GetUserRecordByUserIdAndModuleIdParams) (Userrecord, error)
	ListModules(ctx context.Context, arg ListModulesParams) ([]Module, error)
	ListUserRecords(ctx context.Context, arg ListUserRecordsParams) ([]Userrecord, error)
	UpdateModule(ctx context.Context, arg UpdateModuleParams) error
	UpdateUserRecord(ctx context.Context, arg UpdateUserRecordParams) error
	UpdateUserRecordScore(ctx context.Context, arg UpdateUserRecordScoreParams) error
}

var _ Querier = (*Queries)(nil)
