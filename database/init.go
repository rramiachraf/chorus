package database

import (
	"database/sql"
	"errors"
	"os"
	"path"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Start(f string) error {
	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		return err
	}

	chorusConfigDir := path.Join(userConfigDir, "chorus")
	err = os.Mkdir(chorusConfigDir, 0777)
	if !errors.Is(err, os.ErrExist) && err != nil {
		return err
	}

	err = os.Mkdir(path.Join(chorusConfigDir, "pictures"), 0777)
	if !errors.Is(err, os.ErrExist) && err != nil {
		return err
	}

	dbPath := path.Join(chorusConfigDir, f)
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return err
	}

	DB = db

	err = createTable()
	if err != nil {
		return err
	}

	return nil
}

func createTable() error {
	q := `
	CREATE TABLE IF NOT EXISTS albums (
		id INTEGER PRIMARY KEY,
		name TEXT NOT NULL,
		artist TEXT,
		picture INTEGER REFERENCES pictures(id),
		UNIQUE(name, artist)
	);

	CREATE TABLE IF NOT EXISTS songs (
		id INTEGER PRIMARY KEY,
		title TEXT NOT NULL,
		picture INTEGER REFERENCES pictures(id),
		path TEXT NOT NULL,
		album INTEGER REFERENCES albums(id),
		artist INTEGER REFERENCES artist(id),
		track INTEGER,
		disc INTEGER,
		year INTEGER,
		duration INTEGER,
		mime TEXT
	);

	CREATE TABLE IF NOT EXISTS artists (
		id INTEGER PRIMARY KEY,
		name TEXT NOT NULL UNIQUE
	);

	CREATE TABLE IF NOT EXISTS pictures (
		id INTEGER PRIMARY KEY,
		path TEXT NOT NULL UNIQUE,
		mime TEXT
	)
	`

	_, err := DB.Exec(q)
	if err != nil {
		return err
	}

	return nil
}
