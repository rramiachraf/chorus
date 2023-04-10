package metadata

import (
	"context"
	"database/sql"
	"fmt"
	"io/fs"
	"log"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/rramiachraf/chorus/database"
)

func appendSongs(songs *[]string, p string, isDir bool) {
	isHidden := strings.HasPrefix(p, ".")
	isAudio, _ := regexp.Match(`\.(mp3|flac|wav|m4a)$`, []byte(p))

	if !isDir && isAudio && !isHidden {
		*songs = append(*songs, p)
	}
}

func AddSongsToDB(dirs []string) error {
	var songs []string
	tx, err := database.DB.BeginTx(context.Background(), &sql.TxOptions{})
	if err != nil {
		return err
	}

	for _, dir := range dirs {
		filepath.WalkDir(dir, func(p string, d fs.DirEntry, err error) error {
			if err == nil {
				appendSongs(&songs, p, d.IsDir())
			}
			return nil
		})
	}

	fmt.Printf("Scanning %d songs...\n", len(songs))

	for _, s := range songs {

		err := handleMetadata(tx, s)
		if err != nil {
			log.Println(err)
			continue
		}
	}

	defer func() {
		tx.Commit()
	}()

	return nil
}
