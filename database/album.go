package database

import "database/sql"

type album struct {
	ID      int         `json:"id,omitempty"`
	Name    string      `json:"name"`
	Artist  string      `json:"artist,omitempty"`
	Picture int         `json:"picture"`
	Songs   []albumSong `json:"songs,omitempty"`
}

type albumSong struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Track int    `json:"track"`
	Disc  int    `json:"disc"`
}

func GetAlbums() ([]album, error) {
	q := `
				SELECT albums.id, albums.name, artists.name, albums.picture 
				FROM albums 
				LEFT JOIN artists 
				ON artists.id = albums.artist 
				ORDER BY artists.name
				`

	r, err := DB.Query(q)
	if err != nil {
		return nil, err
	}

	var albums []album

	for r.Next() {
		var a album
		r.Scan(&a.ID, &a.Name, &a.Artist, &a.Picture)
		albums = append(albums, a)
	}

	return albums, nil
}

func GetAlbum(id int) (album, error) {
	q := `
				SELECT albums.name, artists.name, albums.picture
				FROM albums
				LEFT JOIN artists
				ON artists.id = albums.artist
				WHERE albums.id = ?
				`

	var a album

	stmt, err := DB.Prepare(q)
	if err != nil {
		return a, err
	}

	r := stmt.QueryRow(id)
	r.Scan(&a.Name, &a.Artist, &a.Picture)

	q = `
			SELECT id, title, track, disc
			FROM songs 
			WHERE album = ?
			ORDER BY disc, track
			`

	stmt, err = DB.Prepare(q)
	if err != nil {
		return a, err
	}

	rows, err := stmt.Query(id)
	if err != nil {
		return a, err
	}

	for rows.Next() {
		var s albumSong
		rows.Scan(&s.ID, &s.Title, &s.Track, &s.Disc)
		a.Songs = append(a.Songs, s)
	}

	return a, nil
}

func CreateAlbum(tx *sql.Tx, name string, artist string, picture string) error {
	q := `
			INSERT INTO albums (name, artist, picture)
			VALUES (?, (SELECT id FROM artists WHERE name = ?), (SELECT id FROM pictures WHERE path = ?))
			ON CONFLICT (name, artist) DO NOTHING
			`

	stmt, err := tx.Prepare(q)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(name, artist, picture)
	if err != nil {
		return err
	}

	return nil
}
