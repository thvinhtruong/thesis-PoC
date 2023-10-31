package main

import (
	"client/poctest/api"
	"client/poctest/utils"
	"testing"
)

func BenchmarkGetUserRecordV1_OneSingleUserID_CacheEnable(b *testing.B) {
}

func BenchmarkGetUserRecordV1_OneSingleUserID_CacheDisable(b *testing.B) {
}

func BenchmarkGetUserRecordV1_MultipleUserIDs_CacheEnable(b *testing.B) {
}

func BenchmarkGetUserRecordV1_MultipleUserIDs_CacheDisable(b *testing.B) {
}

func GetUserRecordV1API_TDD(isOnly1Record bool, cacheEnable bool, threshold int) {
	configuration := api.NewAPIConfig(cacheEnable)

	if isOnly1Record {
		// random 1 user id within the test size
		requestedUserId := utils.RandomInt(1, threshold)
		configuration.SetAPIEndpoint(utils.IntToString(requestedUserId))
		err := api.MakeCallToApi(configuration.GetAPIEndpoint(), cacheEnable, isOnly1Record)
		if err != nil {
			panic(err)
		}
	} else {
		// random 10 user ids within the test size
		for i := 0; i < threshold; i++ {
			requestedUserId := utils.RandomInt(1, threshold)
			configuration.SetAPIEndpoint(utils.IntToString(requestedUserId))
			err := api.MakeCallToApi(configuration.GetAPIEndpoint(), cacheEnable, isOnly1Record)
			if err != nil {
				panic(err)
			}
		}
	}
}
