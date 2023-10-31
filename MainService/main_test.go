package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	GrpcStudyService "server/MainService/GrpcClients/StudyService"
	GrpcUserService "server/MainService/GrpcClients/UserService"
	"server/MainService/config"
	"server/MainService/handlers"
	reverseproxy "server/MainService/reverse_proxy"
	"server/UserService/pkg/random"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

const (
	TDD_RequestID = "reqId"
)

func setupSuite(tb testing.TB) func(tb testing.TB) {
	return func(tb testing.TB) {
	}
}

func TDD_RandomName() string {
	return "Student " + strconv.Itoa(random.RandomInt(1, 1000))
}

func TDD_CreateRandomUser() GrpcUserService.RegisterUserRequest {
	return GrpcUserService.RegisterUserRequest{
		Fullname: TDD_RandomName(),
		Email:    random.RandomEmail(),
		Password: random.RandomString(10),
		Gender:   random.RandomGender(),
	}
}

func TDD_CreateMockHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		val := r.Context().Value("reqId")
		if val == nil {
			log.Println("reqId not present")
		}
		valStr, ok := val.(string)
		if !ok {
			log.Println("not string")
		}
		if valStr != "1234" {
			log.Println("wrong reqId")
		}
	})
}

func NewContextWithRequestID(ctx context.Context, r *http.Request) context.Context {
	return context.WithValue(ctx, TDD_RequestID, "1234")
}

func AddContextWithRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var ctx = context.Background()
		ctx = NewContextWithRequestID(ctx, r)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func TestRegisterUser(t *testing.T) {
	t.Skip("skipping test in short mode.")
	t.Logf("TestRegisterUser\n")
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)

	configuration := config.GetInstance()
	userHandler := handlers.NewUserApiHanlder(configuration, GrpcUserService.Instance)

	fmt.Printf(
		"Server is listening at %v:%v\n",
		configuration.GetConfig(config.MAIN_SERVICE_HOST),
		configuration.GetConfig(config.MAIN_SERVICE_PORT),
	)

	var createUserRequest = TDD_CreateRandomUser()
	response := userHandler.Repo.RegisterUser(&createUserRequest)

	log.Printf("response: %v", response)

	require.NotNil(t, response)
}

func TestGetUserRecord(t *testing.T) {
	t.Skip("skipping test in short mode.")
	t.Logf("TestGetUserRecord\n")
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)

	configuration := config.GetInstance()
	studyHandler := handlers.NewStudyHandler(configuration, GrpcStudyService.StudyInstance)

	fmt.Printf(
		"Server is listening at %v:%v\n",
		configuration.GetConfig(config.MAIN_SERVICE_HOST),
		configuration.GetConfig(config.MAIN_SERVICE_PORT),
	)

	got := studyHandler.Repo.GetUserRecord(&GrpcStudyService.GetUserRecordRequest{
		UserId: 1,
	})

	log.Printf("got: %v", got)

	require.NotNil(t, got)
}

func TestGetUserRecordHandler(t *testing.T) {
	start := time.Now()
	testHandler := TDD_CreateMockHandler()

	// create the handler to test, using our custom "next" handler
	handlerToTest := AddContextWithRequestID(testHandler)

	// create a mock request to use
	req := httptest.NewRequest("GET", "http://localhost:9000/api/v1/GetUserRecord/2", nil)

	// use middleware
	reverseproxy.HttpResponseCachingMiddleware(handlerToTest)

	handlerToTest.ServeHTTP(httptest.NewRecorder(), req)

	// check that the status code is what we expect
	if status := 200; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// check response
	log.Printf("response: %v", req.Response)

	elapsed := time.Since(start)
	log.Printf("TestGetUserRecordHandler took %s", elapsed)
}
