package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

const (
	ContentTypeJSON = "application/json"
)

func WriteJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", ContentTypeJSON)
	w.WriteHeader(status)

	e := json.NewEncoder(w)
	e.Encode(data)
}

func HandleError(w http.ResponseWriter, status int, err error, display string) {
	if err != nil {
		log.Println(err)
	}

	WriteJSON(w, status, map[string]string{
		"error": display,
	})
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	HandleError(w, http.StatusNotFound, nil, "endpoint not found")
}
