package main

import (
	"net/http"
)

type PresenceResp struct {
	TimeToExpire int `json:"time_to_expire"`
}

func Presence(w http.ResponseWriter, r *http.Request) {
	resp := PlatformResponse{
		Code:    20001,
		Message: "User presence status has been added or updated",
		Response: PresenceResp{
			TimeToExpire: 43200,
		},
	}
	w.WriteHeader(201)
	DefaultJSONEncoder(w, resp)
}
