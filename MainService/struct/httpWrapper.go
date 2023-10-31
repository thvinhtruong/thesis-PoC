package _struct

import (
	"net/http"
	"time"
)

type HttpResponseWrapper struct {
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

type CustomResponseWriter struct {
	http.ResponseWriter
}
