package handlers

import (
	"net/http"

	"github.com/rramiachraf/chorus/database"
)

type stats struct {
	TotalArtists int `json:"totalArtists"`
	TotalSongs   int `json:"totalSongs"`
	TotalAlbums  int `json:"totalAlbums"`
}

func GetStats(w http.ResponseWriter, r *http.Request) {
	totalArtists := database.CountArtists()
	totalSongs := database.CountSongs()
	totalAlbums := database.CountAlbums()

	response := stats{
		TotalArtists: totalArtists,
		TotalSongs:   totalSongs,
		TotalAlbums:  totalAlbums,
	}

	WriteJSON(w, http.StatusOK, response)
}
