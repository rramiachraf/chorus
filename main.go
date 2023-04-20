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
	flag.IntVar(&port, "port", port, "Specify port")
	v := flag.Bool("version", false, "Print version number")
	clean := flag.Bool("clean", false, "Remove config files")
	apiOnly := flag.Bool("api-only", false, "Expose the api only without the frontend")
	https := flag.Bool("https", false, "Enable HTTPS")
	certFile := flag.String("cert-file", "", "TLS certificate file")
	keyFile := flag.String("key-file", "", "TLS private key")
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

	c := &handlers.ServerConfig{
		ApiOnly: *apiOnly,
		Assets:  assets,
		PORT:    port,
	}

	if *https {
		if *certFile == "" || *keyFile == "" {
			log.Fatalln("-https flag must be used in combination with -cert-file and key-file")
		}

		c.TLS.Enabled = true
		c.TLS.CertFile = *certFile
		c.TLS.KeyFile = *keyFile
	}

	handlers.StartServer(c)
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
