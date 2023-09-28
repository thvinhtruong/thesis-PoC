package db

import (
	"context"
	"database/sql"
)

func TearDownDatabase(ctx context.Context, testDB *sql.DB) {
	_, err := testDB.ExecContext(ctx, "TRUNCATE TABLE Module")
	if err != nil {
		panic(err)
	}

	_, err = testDB.ExecContext(ctx, "TRUNCATE TABLE UserRecord")
	if err != nil {
		panic(err)
	}
}
