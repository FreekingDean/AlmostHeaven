package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Server struct {
	r *mux.Router
}

func NewServer() *Server {
	r := mux.NewRouter()
	r.Methods("GET").PathPrefix("/session/get-login-token").HandlerFunc(GetToken)
	r.Methods("POST").PathPrefix("/session/login").HandlerFunc(Login)
	r.Methods("POST").PathPrefix("/log/v3/collect_logdata").HandlerFunc(Log)
	r.Methods("POST").PathPrefix("/log/collect_errordata").HandlerFunc(Log)
	r.Methods("GET").PathPrefix("/ping").HandlerFunc(Ping)
	r.Methods("GET").PathPrefix("/cdp-user/ping").HandlerFunc(Ping)
	r.Methods("GET").PathPrefix("/cdp-user/version-info").HandlerFunc(VersionInfo)
	return &Server{r: r}
}

func (s *Server) Start() {
	log.Println("Server Started")
	http.ListenAndServe(":8000", handlers.LoggingHandler(os.Stdout, s.r))
	log.Println("Server Stopped")
}

func DefaultJSONEncoder(w http.ResponseWriter, d interface{}) {
	err := json.NewEncoder(w).Encode(d)
	if err != nil {
		log.Println("Error ", err)
	}
}
