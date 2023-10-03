package db

import (
	"context"
	"testing"

	"server/UserService/pkg/random"

	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	ctx := context.Background()
	arg := CreateUserParams{
		Fullname: random.RandomName(true),
		Email:    random.RandomEmail(),
		Gender:   random.RandomGender(),
	}

	account, err := testQueries.CreateUser(ctx, arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)
}
