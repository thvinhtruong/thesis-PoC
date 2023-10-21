package utils

import (
	"encoding/json"
	"strconv"
)

func Convert(o interface{}) string {
	res, err := json.Marshal(o)
	if err != nil {
		return "UNKNOWN EXCEPTION"
	}
	return string(res)
}

// Convert the cache key from uint64 to string.
func KeyAsString(key uint64) string {
	return strconv.FormatUint(key, 36)
}
