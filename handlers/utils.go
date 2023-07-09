package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/rramiachraf/chorus/database"
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

type Data interface {
	[]database.Album | []database.Song| []database.Artist
}

type page[T Data] struct {
	Data T   `json:"data"`
	Next int `json:"next,omitempty"`
}

func paginate[T Data](data T, length int, current int) page[T] {
	var p page[T]
	p.Data = data

	if len(data) == length {
		p.Next = current + 1
	}

	return p
}
