package wrapper

import (
	"net/http"
	"time"
)

// Cached http response object, wrapper of HttpResponse
type Response struct {
	// Cached response value
	Value []byte

	// Cached response header
	Header http.Header

	// Expiration date
	Expiration time.Time

	// Last date the value is being accessed, using LRU
	LastAccess time.Time

	// Times a cached response is accessed, using LFU
	Frequency int
}
