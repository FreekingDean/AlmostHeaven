package main

import (
	"net/http"
)

type NotificationResponse struct {
	Code     int            `json:"code"`
	Response []Notification `json:"response"`
}

type Notification struct {
	Description string `json:"description"`
	ID          int    `json:"id"`
	Name        string `json:"name"`
}

func GetNotification(w http.ResponseWriter, r *http.Request) {
	resp := NotificationResponse{
		Code: 18000,
		Response: []Notification{
			Notification{Description: "Friend system", ID: 16479, Name: "Friend"},
			Notification{Description: "Presence system", ID: 11727, Name: "Presence"},
			Notification{Description: "Matchmaking", ID: 1219, Name: "Matchmaking"},
			Notification{Description: "Virtual Currency", ID: 14301, Name: "VCCS"},
		},
	}
	DefaultJSONEncoder(w, resp)
}
