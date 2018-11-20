package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type LogRequest struct {
	Error string `json:"error"`
}

func Log(w http.ResponseWriter, r *http.Request) {
	l := &LogRequest{}
	err := json.NewDecoder(r.Body).Decode(&l)
	log.Println(l.Error)
	resp := buildPlatformSuccess(nil)
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		log.Println(err)
	}
}
