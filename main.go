package main

import (
	"github.com/satori/go.uuid"
)

type GenericPlatformResponse struct {
	PlatformResponse *PlatformResponse `json:"platform_response"`
}

type PlatformResponse struct {
	Code     int         `json:"code"`
	Message  string      `json:"message"`
	Response interface{} `json:"response"`
}

func main() {
	server := NewServer()
	server.Start()
}

func generateUUID() string {
	return uuid.Must(uuid.NewV4()).String()
}

func buildPlatformSuccess(resp interface{}) *GenericPlatformResponse {
	return &GenericPlatformResponse{
		PlatformResponse: &PlatformResponse{
			Code:     2000,
			Message:  "success",
			Response: resp,
		},
	}
}
