package main

import (
	"net/http"
)

func LoginQueueTicket(w http.ResponseWriter, r *http.Request) {
	resp := struct {
		PollInterval  int    `json:"poll_interval"`
		QueuePosition int    `json:"queue_pos"`
		TicketToken   string `json:"ticket_token"`
	}{
		1000000000,
		0,
		"Some-Ticket",
	}
	DefaultJSONEncoder(w, resp)
}
