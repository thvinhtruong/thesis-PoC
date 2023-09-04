package db

import (
	"context"
	"database/sql"
	"fmt"
)

type TxStore struct {
	*Queries
	db *sql.DB
}

func (t *TxStore) Rollback() error {
	return t.Queries.db.(*sql.Tx).Rollback()
}

func (t *TxStore) Commit() error {
	return t.Queries.db.(*sql.Tx).Commit()
}

// a context and a callback function as input, start a transaction
// pass the transaction to the callback function.
// if the callback function returns an error, rollback the transaction
// returns a transaction object or an error
func (store *TxStore) enableTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}
