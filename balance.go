package main

import (
	"net/http"
)

type Bal struct {
	Balance            int    `json:"balance"`
	CurrencyExternalID string `json:"currency_external_id"`
	CurrencyID         int    `json:"currency_id"`
}

func GetBalance(w http.ResponseWriter, r *http.Request) {
	resp := buildPlatformSuccess(
		Bal{
			Balance:            1000,
			CurrencyExternalID: "project76_atoms",
			CurrencyID:         6,
		},
	)
	DefaultJSONEncoder(w, resp)
}

func EntitlementSearch(w http.ResponseWriter, r *http.Request) {
	DefaultJSONEncoder(w, []string{})
}
