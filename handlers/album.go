package handlers

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/rramiachraf/chorus/database"
)

func GetAlbums(w http.ResponseWriter, r *http.Request) {
	a, err := database.GetAlbums()
	if err != nil {
		HandleError(w, http.StatusInternalServerError, err, "error fetching albums")
		return
	}

	WriteJSON(w, http.StatusOK, a)
}

func GetAlbum(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		HandleError(w, http.StatusBadRequest, err, "invalid album id")
		return
	}

	a, err := database.GetAlbum(id)
	if err != nil {
		HandleError(w, http.StatusNotFound, err, "album not found")
		return
	}

	WriteJSON(w, http.StatusOK, a)
}
