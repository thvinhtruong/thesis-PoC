package db

import (
	"context"
	"database/sql"
)

type UserRepository interface {
	RegisterUser(ctx context.Context, arg RegisterUserParams) (RegisterUserResult, error)
	LoginUser(ctx context.Context, arg LoginUserParams) (LoginUserResult, error)
}

func NewRepository(db *sql.DB) UserRepository {
	return &TxStore{
		Queries: New(db),
		db:      db,
	}
}
