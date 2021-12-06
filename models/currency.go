package models

import (
	"encoding/json"
	"net/http"
)

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, _ = w.Write(response)
}

type CurrencyModel struct {
	Currency string  `json:currency`
	Data     float64 `json:data`
}

type ProviderResponse struct {
	Eur float64 `json:"EUR"`
	Usd float64 `json:"USD"`
	Gbp float64 `json:"GBP"`
}
