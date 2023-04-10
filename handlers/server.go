package handlers

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/go-chi/chi"
)

func StartServer(port int) {
	r := chi.NewRouter()

	r.Get("/songs", GetSongs)
	r.Get("/albums", GetAlbums)
	r.Get("/artists", GetArtists)
	r.Get("/listen/{id}", ListenToSong)
	r.Get("/picture/{id}", GetPicture)
	r.Get("/song/{id}", GetSong)
	r.Get("/album/{id}", GetAlbum)
	r.Get("/artist/{id}", GetArtist)
	r.Get("/random/song", GetRandomSong)
	r.Get("/stats", GetStats)
	r.NotFound(http.HandlerFunc(NotFoundHandler))

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
