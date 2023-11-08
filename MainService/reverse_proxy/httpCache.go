package reverseproxy

import (
	"log"
	"net/http"
	"net/http/httptest"
	"server/MainService/reverse_proxy/cacheService/memoryCache"
	_struct "server/MainService/struct"
	"server/MainService/utils"
	"time"
)

var Instane *ClientCachedData

func init() {
	Instane = GetClientInstance()
}

func GetClientInstance() *ClientCachedData {
	if Instane == nil {
		return &ClientCachedData{
			CacheService: memoryCache.NewCacheServiceAdapter(),
		}
	}

	return Instane
}

// treat http response cache as reverse proxy to filter request used by middleware
func HttpResponseCachingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		customWriter := &_struct.CustomResponseWriter{ResponseWriter: w}

		if r.Method != "GET" {
			// Call the next handler
			next.ServeHTTP(w, r)
			return
		}

		// Sort url
		utils.SortURLParams(r.URL)

		// Generate key
		key := utils.GenerateKey(r.URL.String())

		// Check if the response is already cached
		cachedResponse, found := GetClientInstance().CacheService.Get(key)
		if found {
			response := BytesToResponse(cachedResponse)
			response.LastAccess = time.Now()
			response.Frequency += 1

			log.Println("Cached hit!")

			for key, value := range response.Header {
				customWriter.Header()[key] = value
			}

			// if cache found, write the cached value, and no proceed to handler
			w.WriteHeader(http.StatusOK)
			w.Write(response.Value)

		} else {
			customWriter.Header().Set("Content-Type", "application/json")
			customWriter.Header().Set("Cache-Control", "max-age=3600")
			customWriter.Header().Set("Expires", time.Now().Add(time.Hour).Format(http.TimeFormat))
			customWriter.Header().Set("Last-Modified", time.Now().Format(http.TimeFormat))
			customWriter.Header().Set("ETag", "123")

			rec := httptest.NewRecorder()

			httpResponseWrapper := &_struct.HttpResponseWrapper{
				Value:      rec.Body.Bytes(),
				Header:     customWriter.Header(),
				Expiration: time.Now().Add(time.Hour),
				LastAccess: time.Now(),
				Frequency:  1,
			}

			// Set custom writer to cache
			GetClientInstance().CacheService.Set(key, ResponseToBytes(httpResponseWrapper))
		}

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}

func HttpCacheWriter(w http.ResponseWriter, r *http.Request, responseValue []byte, isEnable bool) bool {
	customWriter := &_struct.CustomResponseWriter{ResponseWriter: w}

	if r.Method != "GET" {
		return false
	}

	// Sort url
	utils.SortURLParams(r.URL)

	// Generate key
	key := utils.GenerateKey(r.URL.String())

	// Check if the response is already cached
	cachedResponse, found := GetClientInstance().CacheService.Get(key)
	if found && isEnable {
		response := BytesToResponse(cachedResponse)
		response.LastAccess = time.Now()
		response.Frequency += 1

		for key, value := range response.Header {
			customWriter.Header()[key] = value
		}

		// if cache found, write the cached value, and no proceed to handler
		w.WriteHeader(http.StatusOK)
		w.Write(response.Value)
		return true
	}

	if !found && isEnable && responseValue != nil {
		customWriter.Header().Set("Content-Type", "application/json")
		customWriter.Header().Set("Cache-Control", "max-age=3600")
		customWriter.Header().Set("Expires", time.Now().Add(time.Hour).Format(http.TimeFormat))
		customWriter.Header().Set("Last-Modified", time.Now().Format(http.TimeFormat))
		customWriter.Header().Set("ETag", "123")

		httpResponseWrapper := &_struct.HttpResponseWrapper{
			Value:      responseValue,
			Header:     customWriter.Header(),
			Expiration: time.Now().Add(time.Hour),
			LastAccess: time.Now(),
			Frequency:  1,
		}

		// Set custom writer to cache
		GetClientInstance().CacheService.Set(key, ResponseToBytes(httpResponseWrapper))
	}

	return false
}
