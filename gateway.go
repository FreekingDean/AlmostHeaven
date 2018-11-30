package main

import (
	"net/http"
)

func GatewayLogin(w http.ResponseWriter, r *http.Request) {
	resp := struct {
		ID        int    `json:"id"`
		SessionID string `json:"session_id"`
		Token     string `json:"token"`
		Username  string `json:"username"`
	}{
		1000000000,
		generateUUID(),
		"Some-Gateway-Ticket",
		"SomeUsername",
	}
	DefaultJSONEncoder(w, resp)
}
