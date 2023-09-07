package response

import (
	"encoding/json"
	"log"
	"net/http"
)

func WithJson(w http.ResponseWriter, payload interface{}, code int) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("failed to marshal JSON repsonse %v\n", payload)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}

func WithErr(w http.ResponseWriter, msg string, code int) {
	if code > 499 {
		log.Printf("responding with error code %d, %s\n", code, msg)
	}

	http.Error(w, msg, code)
}
