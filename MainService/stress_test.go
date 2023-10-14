package main

import (
	"log"
	GrpcStudyService "server/MainService/GrpcClients/StudyService"
	"server/MainService/config"
	"server/MainService/handlers"
	"testing"

	"github.com/stretchr/testify/require"
)

func BenchmarkGetUserRecord_Concurrent_OneUserId(t *testing.B) {
	configuration := config.GetInstance()
	studyHandler := handlers.NewStudyHandler(configuration, GrpcStudyService.StudyInstance)

	// wait for 100 concurrent request for user id 1
	ch := make(chan bool, 1000)

	// 100 concurrent request for user id 1
	for i := 0; i < 1000; i++ {
		go func() {
			got := studyHandler.Repo.GetUserRecord(&GrpcStudyService.GetUserRecordRequest{
				UserId: 1,
			})

			if got == nil {
				log.Println("got is nil")
				ch <- false
			} else {
				ch <- true
			}
		}()
	}

	// wait for all request to finish
	for i := 0; i < 100; i++ {
		result := <-ch
		require.True(t, result)
	}
}
