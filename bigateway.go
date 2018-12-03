package main

import (
	"fmt"
	"net/http"
)

func BiGatewayLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `{"level":9,"token":"some-token"}`)
}
