package reverseproxy

import "time"

// Abstraction for cache service which can be used by key/value memory cache or redis
type ICacheService interface {
	// Get retrieves the cached response by a given key
	// if true, cache found and false, otherwise.
	Get(key uint64) ([]byte, bool)

	// Save to cache for a given key until an expiration date.
	Set(key uint64, response []byte, expiration time.Time)

	// Delete value of a given key
	Delete(key uint64)

	// Invalid cache
	Release()
}
