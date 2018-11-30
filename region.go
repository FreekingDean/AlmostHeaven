package main

import (
	"fmt"
	"net/http"
)

func RegionPing(w http.ResponseWriter, _ *http.Request) {
	fmt.Println("GOT HERE")
	w.WriteHeader(200)
}
