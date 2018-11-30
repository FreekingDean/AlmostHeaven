package main

import (
	"net/http"
)

func RegionPing(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(200)
}
