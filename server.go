package main

import (
	"encoding/json"
	"fmt"
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
	r.Methods("POST").PathPrefix("/session/logout").HandlerFunc(Logout)
	r.Methods("POST").PathPrefix("/log/v3/collect_logdata").HandlerFunc(Log)
	r.Methods("POST").PathPrefix("/log/collect_errordata").HandlerFunc(Log)
	r.Methods("POST").PathPrefix("/bigateway/bps/pub/bi/session/event/bulk").HandlerFunc(Log)
	r.Methods("GET").PathPrefix("/ping").HandlerFunc(Ping)
	r.Methods("POST").PathPrefix("/cms/message").HandlerFunc(Message)
	r.Methods("GET").PathPrefix("/cdp-user/ping").HandlerFunc(Ping)
	r.Methods("GET").PathPrefix("/cdp-user/version-info").HandlerFunc(VersionInfo)
	r.Methods("GET").PathPrefix("/titlestorage/v1/products/my-product/platforms/pc/slots/1/branches/prodpc01").HandlerFunc(TitleStorage)
	r.Methods("GET").PathPrefix("/titlestorage/actualstorage/gateways").HandlerFunc(TitleStorageGateways)
	r.Methods("GET").PathPrefix("/notification/v1/system").HandlerFunc(GetNotification)
	r.Methods("HEAD").PathPrefix("/regions/{id}/ping").HandlerFunc(RegionPing)
	r.Methods("POST").PathPrefix("/loginqueue/bps/pub/ticket").HandlerFunc(LoginQueueTicket)
	r.Methods("PUT").PathPrefix("/loginqueue/bps/pub/ticket").HandlerFunc(LoginQueueTicket)
	r.Methods("POST").PathPrefix("/gateway/bps/pub/v2/login").HandlerFunc(GatewayLogin)
	r.Methods("POST").PathPrefix("/gateway/bps/pub/v3/lobby/reconnect").HandlerFunc(GatewayLobbyReconnect)
	r.Methods("POST").PathPrefix("/bigateway/bps/pub/bi/session/create").HandlerFunc(BiGatewayLogin)
	r.Methods("GET").PathPrefix("/social/v3/friends").HandlerFunc(GetFriends)
	r.Methods("POST").PathPrefix("/presence/v1/presence").HandlerFunc(Presence)
	r.Methods("GET").PathPrefix("/pushy/bps/pub/wsa").HandlerFunc(PushySocket)
	r.Methods("GET").PathPrefix("/gateway/bps/pub/character/list").HandlerFunc(CharacterList)
	r.Methods("GET").PathPrefix("/vccs/wallet/balance").HandlerFunc(GetBalance)
	r.Methods("GET").PathPrefix("/entitlements/search").HandlerFunc(EntitlementSearch)
	return &Server{r: r}
}

func (s *Server) Start() {
	log.Println("Server Started")
	http.ListenAndServe(":8000", handlers.LoggingHandler(os.Stdout, s.r))
	log.Println("Server Stopped")
}

func DefaultJSONEncoder(w http.ResponseWriter, d interface{}) {
	output, err := json.Marshal(d)
	if err != nil {
		log.Println("Error ", err)
	}
	fmt.Println(string(output))
	fmt.Fprintf(w, string(output))
}
