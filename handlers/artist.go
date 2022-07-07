package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rramiachraf/chorus/database"
)

func GetArtists(w http.ResponseWriter, r *http.Request) {
	a, err := database.GetArtists()
	if err != nil {
		HandleError(w, http.StatusInternalServerError, err, "error fetching artists")
		return
	}

	WriteJSON(w, http.StatusOK, a)
}

func GetArtist(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
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
