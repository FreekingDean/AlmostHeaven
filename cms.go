package main

import (
	"net/http"
)

type MessageResponse struct {
	CAuthor string `json:"cauthor"`
	CTime   string `json:"ctime"`
	UAuthor string `json:"uauthor"`
	UTime   string `json:"utime"`

	Title   string `json:"title"`
	Content string `json:"content"`

	Language   string `json:"lang"`
	MessageID  int    `json:"message_id"`
	PlatformID int    `json:"platform_id"`
	ProductID  int    `json:"product_id"`
	Public     int    `json:"public"`
	TypeID     int    `json:"type_id"`
	Version    string `json:"version"`
}

func Message(w http.ResponseWriter, _ *http.Request) {
	resp := []*MessageResponse{
		&MessageResponse{
			CAuthor: "Contentful",
			CTime:   "2018-11-13T11:19:27.839000+00:00",
			UAuthor: "Contentful",
			UTime:   "2018-11-13T11:19:27.839000+00:00",

			Content: `{"pre-login":"Hello! FO 78 yo!","main-menu":"FO78"}`,
			Title:   "Fallout 78 Lunch",

			Language:   "en",
			PlatformID: 0,
			MessageID:  1,
			ProductID:  10,
			Public:     1,
			TypeID:     1,
			Version:    "10",
		},
	}
	fullResp := buildPlatformSuccess(resp)
	DefaultJSONEncoder(w, fullResp)
}
