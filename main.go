package main

import (
	"embed"
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"runtime"
	"strings"

	"github.com/rramiachraf/chorus/database"
	"github.com/rramiachraf/chorus/handlers"
	"github.com/rramiachraf/chorus/metadata"
)

var (
	Name       = "chorus"
	Version    = "dev"
	port       = 3000
	dbFileName = "db.sqlite3"
	//go:embed view/dist
	assets embed.FS
)

func main() {
	flag.IntVar(&port, "port", port, "specify port")
	v := flag.Bool("version", false, "print version number")
	clean := flag.Bool("clean", false, "remove config files")
	apiOnly := flag.Bool("api-only", false, "expose the api only without the frontend")
	flag.Parse()

	if *v {
		fmt.Printf("%s %s %s/%s\n", Name, Version, runtime.GOOS, runtime.GOARCH)
		os.Exit(0)
	}

	if *clean {
		err := removeConfigs()
		if err != nil {
			log.Fatalln(err)
		}

		os.Exit(0)
	}

	err := database.Start(dbFileName)
	if err != nil {
		log.Fatalln(err)
	}

	defer database.DB.Close()

	if database.CountSongs() == 0 {
		dirsArg := flag.Arg(0)
		if dirsArg == "" {
			log.Fatalln("music library must be provided to scan")
		}

		dirs := strings.Split(dirsArg, ",")

		for _, d := range dirs {
			_, err := os.Stat(d)
			if err != nil {
				log.Fatalf("`%s` does not exist\n", d)
			}
		}

		if err := metadata.AddSongsToDB(dirs); err != nil {
			log.Fatalln(err)
		}
	}

	handlers.StartServer(*apiOnly, assets, port)
}

func removeConfigs() error {
	c, err := os.UserConfigDir()
	if err != nil {
		return err
	}

	err = os.RemoveAll(path.Join(c, Name))
	if err != nil {
		return err
	}

	return nil
}
