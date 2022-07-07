package metadata

import (
	"database/sql"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"regexp"

	"github.com/dhowden/tag"
	"github.com/rramiachraf/chorus/database"
)

func removeIllegalChars(str string) string {
	rgx := regexp.MustCompile(`[^a-zA-Z_-]`)
	cleaned := rgx.ReplaceAll([]byte(str), []byte(""))
	return string(cleaned)
}

func getPicturesDIR() string {
	configDIR, _ := os.UserConfigDir()
	return path.Join(configDIR, "chorus", "pictures")
}

func savePicture(tx *sql.Tx, picture *tag.Picture, artist string, album string) (string, error) {
	if picture == nil {
		return "", nil
	}

	data := picture.Data
	ext := picture.Ext
	mime := picture.MIMEType

	filename := fmt.Sprintf("%s-%s.%s", removeIllegalChars(artist), removeIllegalChars(album), ext)

	p := path.Join(getPicturesDIR(), filename)
	abs, _ := filepath.Abs(p)

	if err := os.WriteFile(p, data, 0777); err != nil {
		return "", err
	}

	if err := database.CreatePicture(tx, abs, mime); err != nil {
		return "", err
	}

	return abs, nil
}
