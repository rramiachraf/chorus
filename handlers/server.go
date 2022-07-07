package handlers

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func StartServer(port int) {
	r := mux.NewRouter()

	// Development specific
	r.Use(WithCORS)

	r.HandleFunc("/songs", GetSongs).Methods("GET")
	r.HandleFunc("/albums", GetAlbums).Methods("GET")
	r.HandleFunc("/artists", GetArtists).Methods("GET")
	r.HandleFunc("/listen/{id}", ListenToSong).Methods("GET")
	r.HandleFunc("/picture/{id}", GetPicture).Methods("GET")
	r.HandleFunc("/song/{id}", GetSong).Methods("GET")
	r.HandleFunc("/album/{id}", GetAlbum).Methods("GET")
	r.HandleFunc("/artist/{id}", GetArtist).Methods("GET")
	r.HandleFunc("/random/song", GetRandomSong).Methods("GET")
	r.HandleFunc("/stats", GetStats).Methods("GET")
	r.NotFoundHandler = http.HandlerFunc(NotFoundHandler)

	srv := &http.Server{
		Handler:      r,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	portStr := fmt.Sprintf(":%d", port)

	l, err := net.Listen("tcp", portStr)

	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Server is listening on %s", portStr)

	log.Fatalln(srv.Serve(l))
}
