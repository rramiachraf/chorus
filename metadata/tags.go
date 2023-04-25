package metadata

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/dhowden/tag"
	"github.com/rramiachraf/chorus/database"
)

func getMimeType(p string) string {
	ext := path.Ext(p)
	mime := fmt.Sprintf("audio/%s", strings.Replace(ext, ".", "", 1))
	return mime
}

func handleMetadata(tx *sql.Tx, p string) error {
	f, err := os.Open(p)
	if err != nil {
		return err
	}

	defer f.Close()

	var song database.Song

	metadata, err := tag.ReadFrom(f)

	song.Mime = getMimeType(p)
	song.Path = p

	if err != nil {
		song.Title = path.Base(p)
		database.CreateSong(tx, &song, "")
		return nil
	}

	artist := metadata.AlbumArtist()
	title := metadata.Title()
	album := metadata.Album()
	track, _ := metadata.Track()
	year := metadata.Year()
	picture := metadata.Picture()
	disc, _ := metadata.Disc()

	if artist == "" {
		artist = metadata.Artist()
	}

	if title == "" {
		base := path.Base(p)
		title = strings.Replace(base, path.Ext(base), "", 1)
	}

	picturePath, err := savePicture(tx, picture, artist, album)

	if err := database.CreateArtist(tx, artist); err != nil {
		log.Println(err)
	}

	if album != "" {
		err := database.CreateAlbum(tx, album, artist, picturePath)
		if err != nil {
			log.Println(err)
		}
		song.Album = album
	}

	song.Title = title
	song.Artist = artist
	song.Year = year
	song.Track = track
	song.Disc = disc

	if err := database.CreateSong(tx, &song, picturePath); err != nil {
		log.Println(err)
	}

	return nil
}
