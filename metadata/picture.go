package metadata

import (
	"database/sql"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"regexp"

	"github.com/davidbyttow/govips/v2/vips"
	"github.com/dhowden/tag"
	"github.com/rramiachraf/chorus/database"
	"github.com/rramiachraf/chorus/utils"
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

	filename := fmt.Sprintf("%s-%s.webp", removeIllegalChars(artist), removeIllegalChars(album))

	p := path.Join(getPicturesDIR(), filename)
	abs, _ := filepath.Abs(p)

	img, err := vips.NewThumbnailFromBuffer(picture.Data, 350, 350, vips.InterestingAll)
	utils.LogError(err)

	err = img.RemoveMetadata()
	utils.LogError(err)

	webp, _, err := img.ExportWebp(vips.NewWebpExportParams())
	utils.LogError(err)

	err = os.WriteFile(p, webp, 0777)
	utils.LogError(err)

	err = database.CreatePicture(tx, abs, "image/webp")
	utils.LogError(err)

	return abs, nil
}
