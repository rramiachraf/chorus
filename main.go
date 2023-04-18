package main

import (
	"database/sql"
	"embed"
	"flag"
	"log"
	"os"
	"path"
	"strings"

	"github.com/rramiachraf/chorus/database"
	"github.com/rramiachraf/chorus/handlers"
	"github.com/rramiachraf/chorus/metadata"
)

const PORT = 3000

var DB *sql.DB

//go:embed view/dist
var assets embed.FS

func main() {
	// Cleanup
	c, err := os.UserConfigDir()
	if err != nil {
		log.Fatalln(err)
	}

	err = os.RemoveAll(path.Join(c, "chorus"))
	if err != nil {
		log.Fatalln(err)
	}

	// Init
	err = database.Start()
	if err != nil {
		log.Println(err)
	}

	DB = database.DB
	defer DB.Close()

	// Options
	flag.Parse()
	dirsArg := flag.Arg(0)
	dirs := strings.Split(dirsArg, ",")
	if dirsArg == "" {
		log.Fatalln("music library must be provided to scan")
	}

	for _, d := range dirs {
		_, err := os.Stat(d)
		if err != nil {
			log.Fatalf("`%s` does not exist\n", d)
		}
	}

	// Start
	if err := metadata.AddSongsToDB(dirs); err != nil {
		log.Fatalln(err)
	}

	handlers.StartServer(assets, PORT)
}
