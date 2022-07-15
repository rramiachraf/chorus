package database

import "database/sql"

type artist struct {
	ID      int     `json:"id,omitempty"`
	Name    string  `json:"name"`
	Albums  []album `json:"albums,omitempty"`
	Singles []Song  `json:"singles,omitempty"`
}

func GetArtists() ([]artist, error) {
	q := "SELECT id, name FROM artists ORDER BY name"

	r, err := DB.Query(q)
	if err != nil {
		return nil, err
	}

	var artists []artist
	for r.Next() {
		var a artist
		r.Scan(&a.ID, &a.Name)
		artists = append(artists, a)
	}

	return artists, nil
}

func GetArtist(id int) (artist, error) {
	var a artist

	q := "SELECT name FROM artists WHERE id = ?"
	stmt, err := DB.Prepare(q)
	if err != nil {
		return a, err
	}

	r := stmt.QueryRow(id)
	r.Scan(&a.Name)

	q = "SELECT id, name, picture FROM albums WHERE artist = ?"
	stmt, err = DB.Prepare(q)
	if err != nil {
		return a, err
	}

	rows, err := stmt.Query(id)
	if err != nil {
		return a, err
	}

	for rows.Next() {
		var ab album
		rows.Scan(&ab.ID, &ab.Name, &ab.Picture)
		a.Albums = append(a.Albums, ab)
	}

	q = "SELECT id, title, picture FROM songs WHERE artist = ? AND album IS NULL"
	stmt, err = DB.Prepare(q)
	if err != nil {
		return a, err
	}

	rows, err = stmt.Query(id)
	if err != nil {
		return a, err
	}

	for rows.Next() {
		var s Song
		rows.Scan(&s.ID, &s.Title, &s.Picture)
		a.Singles = append(a.Singles, s)
	}

	return a, nil
}

func CreateArtist(tx *sql.Tx, name string) error {
	q := "INSERT INTO artists (name) VALUES(?) ON CONFLICT (name) DO NOTHING"
	stmt, err := tx.Prepare(q)
	if err != nil {
		return err
	}

	if _, err := stmt.Exec(name); err != nil {
		return err
	}

	return nil
}
