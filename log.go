package main

import (
	"net/http"
)

func Log(w http.ResponseWriter, _ *http.Request) {
	resp := buildPlatformSuccess(nil)
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		log.Println(err)
	}
}
