package database

import "database/sql"

type Album struct {
	ID      int         `json:"id,omitempty"`
	Name    string      `json:"name"`
	Artist  string      `json:"artist,omitempty"`
	Picture int         `json:"picture"`
	Year    int         `json:"year,omitempty"`
	Songs   []albumSong `json:"songs,omitempty"`
}

type albumSong struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Track int    `json:"track"`
	Disc  int    `json:"disc"`
}

func GetAlbums(page int) ([]Album, error) {
	q := `
				SELECT albums.id, albums.name, artists.name, albums.picture 
				FROM albums 
				LEFT JOIN artists 
				ON artists.id = albums.artist 
				ORDER BY artists.name
				LIMIT 20 
				OFFSET ?
				`
	r, err := DB.Query(q, (page-1)*20)
	if err != nil {
		return nil, err
	}

	var albums []Album

	for r.Next() {
		var a Album
		r.Scan(&a.ID, &a.Name, &a.Artist, &a.Picture)
		albums = append(albums, a)
	}

	return albums, nil
}

func GetAlbum(id int) (Album, error) {
	q := `
				SELECT albums.name, artists.name, albums.picture, albums.year
				FROM albums
				LEFT JOIN artists
				ON artists.id = albums.artist
				WHERE albums.id = ?
				`

	var a Album

	stmt, err := DB.Prepare(q)
	if err != nil {
		return a, err
	}

	r := stmt.QueryRow(id)
	r.Scan(&a.Name, &a.Artist, &a.Picture, &a.Year)

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

func CreateAlbum(tx *sql.Tx, name string, artist string, picture string, year int) error {
	q := `
			INSERT INTO albums (name, artist, picture, year)
			VALUES (?, (SELECT id FROM artists WHERE name = ?), (SELECT id FROM pictures WHERE path = ?), ?)
			ON CONFLICT (name, artist) DO NOTHING
			`

	stmt, err := tx.Prepare(q)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(name, artist, picture, year)
	if err != nil {
		return err
	}

	return nil
}
