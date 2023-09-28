package db

import (
	"context"
)

type GetUserRecordParams struct {
	UserId int32 `json:"user_id"`
}

type GetUserRecordResult struct {
	UserRecord []GetUserRecordRow
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
