package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

type statusRecorder struct {
	http.ResponseWriter
	statusCode int
}

func MeasureResponseDuration(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rec := statusRecorder{w, 200}

		next.ServeHTTP(&rec, r)

		duration := time.Since(start)
		route := getRoutePattern(r)

		log.Printf("Request duration for %s: %s", route, duration)
		// go func() {
		// 	saveLog("noCache", fmt.Sprintf("%v", duration))
		// }()
		//SaveMemoryProfiling()
	})
}

func getRoutePattern(r *http.Request) string {
	routePattern := r.URL.Path
	if r.URL.RawQuery != "" {
		routePattern += "?" + r.URL.RawQuery
	}

	return routePattern
}

func saveLog(filename string, content ...string) {
	path := "./resource_allocation/" + filename + ".txt"
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()

	for _, c := range content {
		if _, err := f.WriteString(c); err != nil {
			log.Println(err)
		}
	}
}

func ResourceProfiler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Start CPU profiling
		cpuProfile, err := os.Create("cpu_profile.prof")
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer cpuProfile.Close()

		if err := pprof.StartCPUProfile(cpuProfile); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}

func SaveMemoryProfiling() {
	// Start memory profiling
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	// Print memory statistics
	//allocation := fmt.Sprintf("%v\n", memStats.Alloc)
	heapSys := fmt.Sprintf("%v\n", memStats.HeapSys)

	go func() {
		//saveLog("withCache_mem_1record_totalAllocation", allocation)
		saveLog("withCache_mem_1record_totalHeapSystem", heapSys)
	}()

}
