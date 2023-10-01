package _struct

type ApiMessage struct {
	ErrorCode int32       `json:"ErrorCode"`
	Message   string      `json:"Message"`
	Data      interface{} `json:"Data"`
}
