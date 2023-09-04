package utils

import "encoding/json"

func Convert(o interface{}) string {
	res, err := json.Marshal(o)
	if err != nil {
		return "UNKNOWN EXCEPTION"
	}
	return string(res)
}
