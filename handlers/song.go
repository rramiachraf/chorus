package handlers

import (
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rramiachraf/chorus/database"
)

func GetSongs(w http.ResponseWriter, r *http.Request) {
	s, err := database.GetSongs()
	if err != nil {
		HandleError(w, http.StatusInternalServerError, err, "error while fetching songs")
		return
	}

	WriteJSON(w, http.StatusOK, s)
}

func GetSong(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
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
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		HandleError(w, http.StatusBadRequest, err, "invalid song id")
		return
	}

	path, err := database.GetSongPath(id)
	if err != nil {
		HandleError(w, http.StatusNotFound, err, "song not found")
		return
	}

	f, err := os.Open(path)
	if err != nil {
		HandleError(w, http.StatusInternalServerError, err, "cannot play song")
		return
	}

	defer f.Close()

	stat, _ := f.Stat()
	contentLength := strconv.Itoa(int(stat.Size()))

	w.Header().Add("Content-Length", contentLength)
	w.Header().Add("Accept-Ranges", "bytes")
	io.Copy(w, f)
}

func GetRandomSong(w http.ResponseWriter, r *http.Request) {
	id := database.GetRandomSong()
	WriteJSON(w, http.StatusOK, map[string]int{
		"id": id,
	})
}
