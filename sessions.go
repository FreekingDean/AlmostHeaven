package main

import (
	"encoding/json"
	"github.com/FreekingDean/AlmostHeaven/auth"
	"log"
	"net/http"
	"time"
)

const (
	fallout76ApplicationID = "2edd3c0e-db9d-4256-83bf-7fd064122948"
	defaultSessionType     = "basic"
	defaultRefreshTime     = 3600
	defaultExpiryTime      = 7200
)

var (
	invalidLoginResponse = &PlatformResponse{Code: 14018, Message: "Invalid login credentials"}
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

	RefreshTime   int64 `json:"refresh_time"`
	TimeToRefresh int   `json:"time_to_refresh"`
	Expiration    int64 `json:"exp"`
	TimeToExpire  int   `json:"time_to_expire"`

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
			PlatformResponse: invalidLoginResponse,
		}
		err = json.NewEncoder(w).Encode(resp)
		if err != nil {
			log.Println(err)
		}
		w.WriteHeader(401)
		return
	}

	resp := &LoginResponse{
		ApplicationID:   fallout76ApplicationID,
		BUID:            user.ID,
		MasterAccountID: user.ID,
		Username:        user.Username,
		ExternalAccount: struct{}{},
		RefreshTime:     time.Now().Add(defaultRefreshTime).Unix(),
		TimeToRefresh:   defaultRefreshTime,
		Expiration:      time.Now().Add(defaultExpiryTime).Unix(),
		TimeToExpire:    defaultExpiryTime,

		SessionType:  defaultSessionType,
		SessionToken: user.JWT().String(),
	}
	fullResp := buildPlatformSuccess(resp)
	err = json.NewEncoder(w).Encode(fullResp)
	if err != nil {
		log.Println(err)
	}
}
