package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Server struct {
	r *mux.Router
}

func NewServer() *Server {
	r := mux.NewRouter()
	r.Methods("GET").PathPrefix("/session/get-login-token").HandlerFunc(GetToken)
	r.Methods("POST")
	return &Server{r: r}
}

func (s *Server) Start() {
	log.Println("Server Started")
	http.ListenAndServe(":3000", s.r)
	log.Println("Server Stopped")
}

func DefaultJSONEncoder(w http.ResponseWriter, d interface{}) {
	err := json.NewEncoder(w).Encode(d)
	if err != nil {
		log.Println("Error ", err)
	}
}
