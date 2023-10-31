package api

import (
	"client/poctest/utils"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func MakeCallToApi(endpoint string, cacheEnable bool, onlyOneRequest bool) error {
	startTime := time.Now()
	started := startTime.Format("2006-01-02 15:04:05")
	response, err := http.Get(endpoint)
	if err != nil {
		log.Printf("Error when making call to API: %s", err)
		return err
	}

	latency := time.Since(startTime)

	responseData := fmt.Sprintf("withCache %s %s %s ", response.Request.Method, response.Request.URL.Path, utils.IntToString(response.StatusCode))

	// Save log in file
	// Using different goroutine not to block the main thread
	go func() {
		if cacheEnable && onlyOneRequest {
			saveLog("withCache_OnlyOneRequest", started, responseData, latency.String())
		} else if cacheEnable && !onlyOneRequest {
			saveLog("withCache-MultipleRequest", started, responseData, latency.String())
		} else if !cacheEnable && onlyOneRequest {
			saveLog("noCache_OnlyOneRequest", started, responseData, latency.String())
		} else {
			saveLog("noCache_MultipleRequest", started, responseData, latency.String())
		}
	}()

	return nil
}

func saveLog(name string, s ...string) {
	// open file
	path := "../log" + name + ".log"
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	for _, items := range s {
		if _, err := file.WriteString(items + " "); err != nil {
			panic(err)
		}
	}

	if _, err := file.WriteString("\n"); err != nil {
		log.Fatal(err)
	}
}
