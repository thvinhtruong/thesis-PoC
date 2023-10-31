package main

import (
	"fmt"
	"log"
	"net/http/httptest"
	GrpcStudyService "server/MainService/GrpcClients/StudyService"
	"server/MainService/config"
	"server/MainService/handlers"
	"server/UserService/pkg/random"
	"testing"
)

var table = []struct {
	input int
}{
	{input: 100},
	{input: 1000},
	{input: 74382},
	// {input: 382399},
}

var studyHandler handlers.StudyHandler

func init() {
	// init service
	configuration := config.GetInstance()
	studyHandler = handlers.NewStudyHandler(configuration, GrpcStudyService.StudyInstance)
}

func BenchmarkGetUserRecord_Concurrent_OneUserId(t *testing.B) {
	t.Skip("skipping test in short mode.")
	// wait for 100 concurrent request for user id 1
	ch := make(chan bool, 1000)
	// defer close(ch)

	// 100 concurrent request for user id 1
	for i := 0; i < 1000; i++ {
		go func() {
			got := studyHandler.Repo.GetUserRecord(&GrpcStudyService.GetUserRecordRequest{
				UserId: 1,
			})

			if got == nil {
				ch <- false
			} else {
				ch <- true
			}
		}()
	}

	// wait for all request to finish
	for i := 0; i < 100; i++ {
		result := <-ch
		if !result {
			log.Printf("record number %v is nil", i)
			t.Fail()
			break
		}
	}
}

func BenchmarkGetUserRecord_Concurrent_MultipleUserID(t *testing.B) {
	t.Skip("skipping test in short mode.")
	// different input size
	for _, v := range table {
		t.Run(fmt.Sprintf("input_size_%d", v.input), func(t *testing.B) {
			testSize := v.input

			// wait for 100 concurrent request for user id 1
			ch := make(chan bool, testSize)

			// 100 concurrent request for user id 1
			for i := 0; i < testSize; i++ {
				go func() {
					got := studyHandler.Repo.GetUserRecord(&GrpcStudyService.GetUserRecordRequest{
						UserId: int32(random.RandomInt(1, 100)),
					})

					if got == nil {
						ch <- false
					} else {
						ch <- true
					}
				}()
			}

			// wait for all request to finish
			for i := 0; i < 100; i++ {
				result := <-ch
				if !result {
					log.Printf("record number %v is nil", i)
					t.Fail()
					break
				}
			}
		})
	}
}

func BenchmarkGetUserRecord_Concurrent_MultipleUserID_WithHandler(t *testing.B) {
	//t.Skip("skipping test in short mode.")
	// requestUrl := "http://localhost:9000/api/v1/GetUserRecord/" + fmt.Sprintf("%d", random.RandomInt(1, 100))
	requestUrlWithOneID := "http://localhost:9000/api/v1/GetUserRecord/1"

	for _, v := range table {
		t.Run(fmt.Sprintf("input_size_%d", v.input), func(t *testing.B) {
			testSize := v.input

			// wait for 100 concurrent request for user id 1
			ch := make(chan bool, testSize)

			// 100 concurrent request for user id 1
			for i := 0; i < testSize; i++ {
				go func() {
					isSuccess := TDD_MockHandlerStressTest(requestUrlWithOneID, true)
					ch <- isSuccess
				}()
			}

			// wait for all request to finish
			for i := 0; i < 100; i++ {
				result := <-ch
				if !result {
					t.Fail()
					break
				}
			}
		})
	}

}

func TDD_MockHandlerStressTest(requestUrl string, enableCache bool) bool {
	testHandler := TDD_CreateMockHandler()

	// create the handler to test, using our custom "next" handler
	handlerToTest := AddContextWithRequestID(testHandler)

	// create a mock request to use
	req := httptest.NewRequest("GET", requestUrl, nil)

	// use middleware
	// if enableCache {
	// 	// reverseproxy.HttpResponseCachingMiddleware(handlerToTest)
	// }

	handlerToTest.ServeHTTP(httptest.NewRecorder(), req)

	return req != nil
}
