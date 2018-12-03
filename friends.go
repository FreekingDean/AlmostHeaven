package main

import (
	"net/http"
)

type FriendList struct {
	FriendList []string `json:"FriendList"`
	Count      int      `json:"total_count"`
}

func GetFriends(w http.ResponseWriter, r *http.Request) {
	resp := buildPlatformSuccess(
		FriendList{
			FriendList: []string{},
			Count:      0,
		},
	)
	DefaultJSONEncoder(w, resp)
}
