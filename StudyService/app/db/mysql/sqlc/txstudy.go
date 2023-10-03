package db

import (
	"context"
)

type CreateModuleResult struct {
	ModuleId int32 `json:"module_id"`
}

type CreateUserRecordResult struct {
	UserId int32 `json:"user_id"`
}

type GetUserRecordParams struct {
	UserId int32 `json:"user_id"`
}

type GetUserRecordResult struct {
	UserRecord []GetUserRecordRow
}

func (u *TxStore) CreateModuleTx(ctx context.Context, arg CreateModuleParams) (CreateModuleResult, error) {
	var result CreateModuleResult
	err := u.enableTx(ctx, func(q *Queries) error {
		var err error
		sqlResult, err := q.CreateModule(ctx, arg)
		if err != nil {
			return err
		}

		id, err := sqlResult.LastInsertId()
		if err != nil {
			return err
		}

		result.ModuleId = int32(id)

		return nil
	})

	return result, err
}

func (u *TxStore) CreateUserRecordTx(ctx context.Context, arg CreateUserRecordParams) (CreateUserRecordResult, error) {
	var result CreateUserRecordResult
	err := u.enableTx(ctx, func(q *Queries) error {
		err := q.CreateUserRecord(ctx, arg)
		if err != nil {
			return err
		}

		result.UserId = arg.UserID

		return nil
	})

	return result, err
}

func (u *TxStore) GetUserRecordTx(ctx context.Context, arg GetUserRecordParams) (GetUserRecordResult, error) {
	var result GetUserRecordResult
	err := u.enableTx(ctx, func(q *Queries) error {
		var err error
		userRecord, err := q.GetUserRecord(ctx, arg.UserId)
		if err != nil {
			return err
		}

		result.UserRecord = userRecord

		return nil
	})

	return result, err
}
