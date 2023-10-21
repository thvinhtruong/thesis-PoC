package utils

import (
	"bytes"
	"encoding/gob"
	"log"
	"server/ReverseProxy/wrapper"
	"time"
)

// BytesToResponse converts bytes array into Response data structure.
func BytesToResponse(b []byte) wrapper.Response {
	var r wrapper.Response
	dec := gob.NewDecoder(bytes.NewReader(b))
	dec.Decode(&r)

	return r
}

// LRU algorithm
func LRU(store map[uint64][]byte) {
	selectedKey := uint64(0)
	lastAccess := time.Now()

	var cacheSize int
	var hit bool

	for k, v := range store {
		response := BytesToResponse(v)
		if response.LastAccess.Before(lastAccess) {
			selectedKey = k
			lastAccess = response.LastAccess
			cacheSize, hit = len(v), true
		}
	}

	if hit {
		log.Println("hit", cacheSize)
	}

	delete(store, selectedKey)
}

func LFU(store map[uint64][]byte) {
	selectedKey := uint64(0)
	frequency := 2147483647

	var cacheSize int
	var hit bool

	for k, v := range store {
		response := BytesToResponse(v)
		if response.Frequency < frequency {
			selectedKey = k
			frequency = response.Frequency
			cacheSize, hit = len(v), true
		}
	}

	if hit {
		log.Println("hit", cacheSize)
	}

	delete(store, selectedKey)
}
