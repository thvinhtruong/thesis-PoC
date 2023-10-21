package main

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	cache           = make(map[string]*cachedResponse)
	cacheMutex      sync.RWMutex
	cacheExpiration = time.Minute * 5
)

type cachedResponse struct {
	Response    *http.Response
	LastFetched time.Time
}

func main() {
	// Define target gRPC service address
	targetAddress := "localhost:9000"

	// Create a gRPC connection to the target service
	conn, _ := grpc.Dial(targetAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))

	// Create a reverse proxy
	proxy := &GRPCProxy{targetConn: conn}

	// Serve the reverse proxy
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		proxy.ServeHTTP(w, r)
	})

	// Start the server
	fmt.Println("gRPC reverse proxy with caching is running on :8000")
	http.ListenAndServe(":8000", nil)
}

type GRPCProxy struct {
	targetConn *grpc.ClientConn
}

func (p *GRPCProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Extract gRPC method from the request path
	method := strings.TrimPrefix(r.URL.Path, "/")

	// Check if the response is already cached
	cacheMutex.RLock()
	cachedResp, found := cache[method]
	cacheMutex.RUnlock()

	if found && time.Since(cachedResp.LastFetched) < cacheExpiration {
		fmt.Println("Serving from cache:", method)
		p.writeCachedResponse(w, cachedResp.Response)
		return
	}

	// Fetch the response from the gRPC service
	resp, err := p.fetchGRPCResponse(method)
	if err != nil {
		http.Error(w, "Error fetching gRPC response", http.StatusInternalServerError)
		return
	}

	// Cache the response
	cacheMutex.Lock()

	cache[method] = &cachedResponse{
		Response:    resp,
		LastFetched: time.Now(),
	}
	cacheMutex.Unlock()

	// Listen and Serve

	fmt.Println("Fetched from gRPC:", method)
	p.writeCachedResponse(w, resp)
}

func (p *GRPCProxy) fetchGRPCResponse(method string) (*http.Response, error) {
	// You would implement logic to convert gRPC request to HTTP request here
	// You can use the targetConn to make gRPC requests to the actual service
	// Convert the gRPC response to an HTTP response

	return nil, nil
}

func (p *GRPCProxy) writeCachedResponse(w http.ResponseWriter, resp *http.Response) {
	// Write the cached response to the client's HTTP response
	// You would implement copying response headers, status, and body here
}
