package reverseproxy

import (
	"bytes"
	"encoding/gob"
	"server/MainService/reverse_proxy/cacheService/memoryCache"
	_struct "server/MainService/struct"
	"time"
)

type ClientCachedData struct {
	CacheService             *memoryCache.CacheServiceAdapter
	HttpResponse             _struct.HttpResponseWrapper
	TTL                      time.Duration
	IsWrittenToExpiredHeader bool
}

func BytesToResponse(b []byte) _struct.HttpResponseWrapper {
	var r _struct.HttpResponseWrapper
	dec := gob.NewDecoder(bytes.NewReader(b))
	dec.Decode(&r)

	return r
}

// Bytes converts Response data structure into bytes array.
func ResponseToBytes(response *_struct.HttpResponseWrapper) []byte {
	var b bytes.Buffer
	enc := gob.NewEncoder(&b)
	enc.Encode(&response)

	return b.Bytes()
}
