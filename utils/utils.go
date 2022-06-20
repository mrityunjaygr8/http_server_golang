package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	response, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(response)
}

type errorBody struct {
	Error string `json:"error"`
}

func RespondWithError(w http.ResponseWriter, code int, err error) {
	e := errorBody{Error: err.Error()}
	RespondWithJSON(w, code, e)
}
