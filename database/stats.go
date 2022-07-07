package database

func CountArtists() int {
	q := "SELECT COUNT(name) FROM artists"

	var count int

	r := DB.QueryRow(q)
	r.Scan(&count)

	return count
}

func CountSongs() int {
	q := "SELECT COUNT(id) FROM songs"

	var count int

	r := DB.QueryRow(q)
	r.Scan(&count)

	return count
}

func CountAlbums() int {
	q := "SELECT COUNT(id) FROM albums"

	var count int

	r := DB.QueryRow(q)
	r.Scan(&count)

	return count
}
