package main

import (
	"net/http"
)

type Character struct {
	Created    int    `json:"created"`
	ID         int    `json:"id"`
	Updated    int    `json:"updated"`
	Level      int    `json:"level"`
	Region     int    `json:"region"`
	Name       string `json:"name"`
	IsComplete bool   `json:"is_complete"`
}

func CharacterList(w http.ResponseWriter, r *http.Request) {
	chars := []Character{
		Character{
			Created:    1540949622619,
			ID:         374603123777353700,
			Updated:    1543807903111,
			Level:      16,
			Region:     2,
			Name:       "BananaBoyDean",
			IsComplete: true,
		},
	}
	resp := struct {
		Characters []Character `json:"characters"`
	}{
		Characters: chars,
	}
	DefaultJSONEncoder(w, resp)
}
