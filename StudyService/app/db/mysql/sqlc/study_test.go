package db

// func TestCreateModule(t *testing.T) {
// 	mockModuleData := mock.MockModuleData("FA")
// 	for _, module := range mockModuleData {
// 		arg := CreateModuleParams{
// 			Name:        module,
// 			Description: sql.NullString{String: "test", Valid: true},
// 		}
// 		result, err := testQueries.CreateModule(context.Background(), arg)
// 		require.NoError(t, err)

// 		id, err := result.LastInsertId()

// 		require.NoError(t, err)
// 		require.NotEmpty(t, id)
// 		require.NotZero(t, id)
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
