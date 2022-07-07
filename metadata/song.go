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
	"time"

	"github.com/briandowns/spinner"
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

	sp := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	sp.Start()

	for i, s := range songs {
		sp.Suffix = fmt.Sprintf(" Scanning songs... (%d/%d)", i+1, len(songs))

		err := handleMetadata(tx, s)
		if err != nil {
			log.Println(err)
			continue
		}

		if i+1 == len(songs) {
			sp.Stop()
		}
	}

	defer func() {
		tx.Commit()
	}()

	return nil
}
