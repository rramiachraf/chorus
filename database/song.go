package database

import (
	"database/sql"
)

type Song struct {
	ID      int    `json:"id,omitempty"`
	Title   string `json:"title"`
	Artist  string `json:"artist,omitempty"`
	Album   string `json:"album,omitempty"`
	Path    string `json:"path,omitempty"`
	Track   int    `json:"track,omitempty"`
	Year    int    `json:"year,omitempty"`
	Mime    string `json:"mime,omitempty"`
	Picture int    `json:"picture,omitempty"`
	Disc    int    `json:"disc,omitempty"`
}

func GetSongs(page int) ([]Song, error) {
	q := `
		SELECT 
		songs.id, songs.title, songs.picture, 
		artists.name as artist, 
		UPPER(REPLACE(mime, "audio/", "")) as mime 
		FROM songs
		LEFT JOIN artists ON songs.artist = artists.id
		ORDER BY songs.title
		LIMIT 20
		OFFSET ?
		`

	var songs []Song

	r, err := DB.Query(q, (page-1)*20)
	if err != nil {
		return nil, err
	}

	for r.Next() {
		var s Song
		r.Scan(&s.ID, &s.Title, &s.Picture, &s.Artist, &s.Mime)
		songs = append(songs, s)
	}

	return songs, nil
}

func GetSong(id int) (Song, error) {
	q := `
			SELECT songs.title, artists.name, songs.picture 
			FROM songs
			LEFT JOIN artists ON artists.id = songs.artist
			WHERE songs.id = ?
			`

	var s Song

	stmt, err := DB.Prepare(q)
	if err != nil {
		return s, err
	}

	r := stmt.QueryRow(id)
	r.Scan(&s.Title, &s.Artist, &s.Picture)

	return s, nil
}

func GetSongPath(id int) (string, error) {
	q := "SELECT path FROM songs WHERE id = ?"
	var path string

	stmt, err := DB.Prepare(q)
	if err != nil {
		return path, err
	}

	r := stmt.QueryRow(id)
	r.Scan(&path)

	return path, err
}

func CreateSong(tx *sql.Tx, s *Song, picture string) error {
	q := `
				INSERT INTO songs (title, picture, path, album, artist, track, disc, year, mime)
				VALUES(
					?, 
					(SELECT id FROM pictures WHERE path = ?), 
					?,
					(SELECT id FROM albums WHERE name = ? AND artist = (SELECT id FROM artists WHERE name = ?)),
					(SELECT id FROM artists WHERE name = ?),
					?,
					?,
					?,
					?
				)
			`

	stmt, err := tx.Prepare(q)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(s.Title, picture, s.Path, s.Album, s.Artist, s.Artist, s.Track, s.Disc, s.Year, s.Mime)
	if err != nil {
		return err
	}

	return nil
}

func GetRandomSong() int {
	var id int

	q := "SELECT id FROM songs ORDER BY RANDOM() LIMIT 1"
	r := DB.QueryRow(q)
	r.Scan(&id)

	return id
}
