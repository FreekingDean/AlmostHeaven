package main

import (
	"enncoding/json"
	"fmt"
	"log"
	"net/http"
)

type VersionInfoResponse struct {
	Branch              string   `json:"BRANCH"`
	BuildTime           string   `json:"BUILD_TIME"`
	Commit              string   `json:"COMMIT"`
	Dirty               bool     `json:"DIRTY"`
	MinClientLibVersion []string `json:"MIN_CLIENT_LIB_VERSION"`
	Name                string   `json:"NAME"`
	Tag                 string   `json:"TAG"`
	Version             []int    `json:"VERSION"`
}

func VersionInfo(w http.ResponseWriter, _ *http.Request) {
	resp := &VersionInfoResponse{
		Branch:              "release-1.51",
		BuildTime:           "Thu Sep 27 11:28:59 UTC 2018",
		Commit:              "e1180f512d7c5078404d88a610e5fbccfe645820",
		Dirty:               false,
		MinClientLibVersion: []string{},
		Name:                "",
		Tag:                 "",
		Version:             []int{1, 51, 1},
	}
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		log.Println(err)
	}
}
