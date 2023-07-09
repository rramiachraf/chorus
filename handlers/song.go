package handlers

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/rramiachraf/chorus/database"
)

func GetSongs(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))

	s, err := database.GetSongs(page)
	if err != nil {
		HandleError(w, http.StatusInternalServerError, err, "error while fetching songs")
		return
	}

	WriteJSON(w, http.StatusOK, paginate(s, 20, page))
}

func GetSong(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		HandleError(w, http.StatusBadRequest, err, "invalid song id")
		return
	}

	s, err := database.GetSong(id)
	if err != nil {
		HandleError(w, http.StatusNotFound, err, "song not found")
		return
	}

	WriteJSON(w, http.StatusOK, s)
}

func ListenToSong(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		HandleError(w, http.StatusBadRequest, err, "invalid song id")
		return
	}

	path, err := database.GetSongPath(id)
	if err != nil {
		HandleError(w, http.StatusNotFound, err, "song not found")
		return
	}

	http.ServeFile(w, r, path)
}

func GetRandomSong(w http.ResponseWriter, r *http.Request) {
	id := database.GetRandomSong()
	WriteJSON(w, http.StatusOK, map[string]int{
		"id": id,
	})
}
