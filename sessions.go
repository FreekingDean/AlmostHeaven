package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func GetToken(w http.ResponseWriter, _ *http.Request) {
	resp := &GenericPlatformResponse{
		PlatformResponse: &PlatformResponse{
			Code:     2000,
			Message:  "success",
			Response: generateUUID(),
		},
	}

	DefaultJSONEncoder(w, resp)
}

type LoginResponse struct {
	ApplicationID string `json:"application_id"`

	BUID            string   `json:"buid"` //Guessing this is Bethesda User ID
	MasterAccountID string   `json:"master_account_id"`
	Username        string   `json:"username"`
	ExternalAccount struct{} `json:"external_account"`

	RefreshTime   int `json:"refresh_time"`
	TimeToRefresh int `json:"time_to_refresh"`
	Expiration    int `json:"exp"`
	TimeToExpire  int `json:"time_to_expire"`

	SessionType  string `json:"session_type"`  //Should be "basic"
	SessionToken string `json:"session_token"` //JWT
}

type LoginRequest struct {
	ClientID string `json:"client_id"`
	Language string `json:"language"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println(err)
	}
	user := auth.BasicAuth(req.Username, req.Password)
	if user == nil {
		resp := &GenericPlatformResponse{
			PlatformResponse: &PlatformResponse{
				Code:    14018,
				Message: "Invalid login credentials",
			},
		}
		err = json.NewEncoder(w).Encode()
		if err != nil {
			log.Println(err)
		}
		w.WriteHeader(401)
		return
	}
	fmt.Fprintf(w, user.GetJWT())
}
