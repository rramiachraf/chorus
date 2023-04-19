package handlers

import (
	"embed"
	"fmt"
	"io"
	"log"
	"mime"
	"net"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/go-chi/chi"
)

func fileServer(fs embed.FS) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fName := "view/dist/index.html"
		exts := []string{"js", "css", "png", "jpg", "jpeg", "webp"}

		for _, ext := range exts {
			if strings.HasSuffix(r.URL.Path, ext) {
				fName = fmt.Sprintf("view/dist%s", r.URL.Path)
			}
		}

		f, err := fs.Open(fName)
		if err != nil {
			HandleError(w, http.StatusInternalServerError, err, "something went wrong")
		}

		defer f.Close()

		ext := filepath.Ext(r.URL.Path)
		w.Header().Set("Content-Type", mime.TypeByExtension(ext))

		_, err = io.Copy(w, f)
		if err != nil {
			log.Println(err)
		}
	})
}

func StartServer(apiOnly bool, assets embed.FS, port int) {
	r := chi.NewRouter()

	r.Route("/api", func(r chi.Router) {
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
	})

	if !apiOnly {
		r.Handle("/*", fileServer(assets))
	}

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
