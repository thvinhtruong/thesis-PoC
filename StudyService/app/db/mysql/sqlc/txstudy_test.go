package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetUserRecordTx(t *testing.T) {
	testDB, err := CreateTestDB()
	if err != nil {
		t.Fatal(err)
	}

	repository := NewStudyRepository(testDB)

	userRecord, err := repository.GetUserRecordTx(context.Background(), GetUserRecordParams{
		UserId: 1,
	})

	if err != nil {
		t.Fatal(err)
	}

	require.NotNil(t, userRecord)
	require.NotEmpty(t, userRecord.UserRecord)

	for _, item := range userRecord.UserRecord {
		require.NotEmpty(t, item.Name)
		require.NotEmpty(t, item.Weight)
		require.NotEmpty(t, item.Score)
	}
}
