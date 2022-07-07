package database

import (
	"database/sql"
	"errors"
)

func GetPicturePath(id int) (string, error) {
	var path string

	q := "SELECT path FROM pictures WHERE id = ?"

	stmt, err := DB.Prepare(q)
	if err != nil {
		return path, err
	}

	r := stmt.QueryRow(id)
	r.Scan(&path)

	if path == "" {
		return path, errors.New("picture not found")
	}

	return path, nil
}

func CreatePicture(tx *sql.Tx, path string, mime string) error {
	stmt, err := tx.Prepare("INSERT INTO pictures (path, mime) VALUES (?, ?) ON CONFLICT (path) DO NOTHING")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(path, mime)
	if err != nil {
		return err
	}

	return nil
}
