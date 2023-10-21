package memoryCache

import (
	"sync"
	"time"
)

var (
	cacheExpiration = time.Minute * 5
)

type CacheServiceAdapter struct {
	cacheMutex sync.RWMutex
	storage    map[uint64][]byte
	// cacheState CacheState
}

func (service *CacheServiceAdapter) Get(key uint64) ([]byte, bool) {
	service.cacheMutex.RLock()
	cachedResp, found := service.storage[key]
	service.cacheMutex.RUnlock()

	if found {
		if time.Since(time.Now()) > cacheExpiration {
			// Temporary: do nothing
			RefreshCache()
		}
		return cachedResp, true
	}

	return nil, false
}

func (service *CacheServiceAdapter) Set(key uint64, response []byte) error {
	service.cacheMutex.Lock()
	defer service.cacheMutex.Unlock()

	if _, ok := service.storage[key]; ok {
		// Known key, overwrite previous item.
		service.storage[key] = response
		return nil
	}

	service.storage[key] = response
	return nil
}

func (service *CacheServiceAdapter) Delete(key uint64) {
	service.cacheMutex.Lock()
	defer service.cacheMutex.Unlock()

	// key not found, do nothing.
	if _, ok := service.storage[key]; !ok {
		return
	}

	delete(service.storage, key)
}

func (service *CacheServiceAdapter) Release() {
	service.cacheMutex.Lock()
	service.storage = make(map[uint64][]byte)
	service.cacheMutex.Unlock()
}

// Refresh new data if query expired cache key
func RefreshCache() {
	time.Sleep(cacheExpiration)
}
