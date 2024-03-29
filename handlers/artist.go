package handlers

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/rramiachraf/chorus/database"
)

func GetArtists(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))

	a, err := database.GetArtists(page)
	if err != nil {
		HandleError(w, http.StatusInternalServerError, err, "error fetching artists")
		return
	}

	WriteJSON(w, http.StatusOK, paginate(a, 24, page))
}

func GetArtist(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		HandleError(w, http.StatusBadRequest, nil, "invalid artist id")
		return
	}

	a, err := database.GetArtist(id)
	if err != nil {
		HandleError(w, http.StatusNotFound, err, "error fetching artist")
		return
	}

	WriteJSON(w, http.StatusOK, a)
}
