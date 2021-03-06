package handlers

import (
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rramiachraf/chorus/database"
)

func GetPicture(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		HandleError(w, http.StatusBadRequest, err, "invalid picture id")
		return
	}

	path, err := database.GetPicturePath(id)
	if err != nil {
		HandleError(w, http.StatusNotFound, err, "picture not found")
		return
	}

	f, err := os.Open(path)
	if err != nil {
		HandleError(w, http.StatusInternalServerError, err, "error fetching picture")
		return
	}

	defer f.Close()
	io.Copy(w, f)
}
