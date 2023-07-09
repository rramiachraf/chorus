# Chorus [![badge](https://github.com/rramiachraf/chorus/actions/workflows/main.yml/badge.svg)](https://github.com/rramiachraf/chorus/actions/workflows/main.yml)
Enjoy your music within the browser realm.

![Screenshot](https://raw.githubusercontent.com/rramiachraf/chorus/main/screenshot.png)

## Requirements
- [Libvips 8.10+](https://github.com/libvips/libvips)
- Modern Browser

## Installation
See [releases](https://github.com/rramiachraf/chorus/releases) for pre-built binaries.

### Docker
```bash
git clone https://github.com/rramiachraf/chorus
cd chorus
docker buildx build -t chorus .
docker run -p 3000:3000 -v $MUSIC_DIR:/music chorus
```

### Building from source
#### Linux
Dependencies:

- Go v1.18+
- Node.js v16+
- GCC
- Libvips
 
```bash
git clone https://github.com/rramiachraf/chorus
cd chorus
make build-linux
```

#### Windows
Dependencies:

- Go v1.18+
- Node.js v16+
- Mingw-w64
- Libvips
 
```
git clone https://github.com/rramiachraf/chorus
cd chorus
make build-windows
```

## Usage
```
Usage of chorus:
  -api-only
    	Expose the api only without the frontend
  -cert-file file
    	TLS certificate file
  -clean
    	Remove config files
  -https
    	Enable HTTPS
  -key-file file
    	TLS private key file
  -port number
    	Specify port number (default 3000)
  -version
    	Print version number
```
```bash
chorus $MUSIC_DIR
```
**Note**: You can specify more directories by seperating them using a comma.
e.g. `chorus "$MUSIC_DIR,$MUSIC_DIR2"`.

## Author
[Achraf RRAMI](https://github.com/rramiachraf) (2022 - present)

## License
Apache License 2.0
