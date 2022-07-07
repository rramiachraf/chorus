package main

import (
	"database/sql"
	"flag"
	"log"
	"os"
	"path"
	"strings"

	"github.com/rramiachraf/chorus/database"
	"github.com/rramiachraf/chorus/handlers"
	"github.com/rramiachraf/chorus/metadata"
)

var DB *sql.DB

const PORT = 3000

func main() {
	// Cleanup
	if c, err := os.UserConfigDir(); err == nil {
		os.RemoveAll(path.Join(c, "chorus"))
	}

	// Init
	err := database.Start()
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

	handlers.StartServer(PORT)
}
