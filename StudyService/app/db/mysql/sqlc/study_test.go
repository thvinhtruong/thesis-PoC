package db

// func TestCreateModule(t *testing.T) {
// 	for i := 0; i < 10; i++ {
// 		arg := CreateModuleParams{
// 			Name:        random.RandomTopic(true),
// 			Description: sql.NullString{String: "test", Valid: true},
// 		}

// 		err := testQueries.CreateModule(context.Background(), arg)
// 		require.NoError(t, err)
// 	}
// }

// func TestCreateUserRecord(t *testing.T) {
// 	arg := CreateUserRecordParams{
// 		UserID:   1,
// 		ModuleID: int32(random.RandomInt(1, 10)),
// 		Weight:   int32(random.RandomInt(1, 10)),
// 		Score:    int32(random.RandomInt(1, 10)),
// 	}
// 	err := testQueries.CreateUserRecord(context.Background(), arg)
// 	require.NoError(t, err)
// }

// func TestGetUserRecord(t *testing.T) {
// 	record, err := testQueries.GetUserRecord(context.Background(), 1)
// 	require.NoError(t, err)
// 	require.NotEmpty(t, record)
// 	require.Greater(t, len(record), 1)
// }
