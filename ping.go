package main

import (
	"fmt"
	"net/http"
)

func Ping(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "pong")
}
