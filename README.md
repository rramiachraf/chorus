# Chorus [![badge](https://github.com/rramiachraf/chorus/actions/workflows/main.yml/badge.svg)](https://github.com/rramiachraf/chorus/actions/workflows/main.yml)
Enjoy your music within the browser realm.

![Screenshot](https://raw.githubusercontent.com/rramiachraf/chorus/main/screenshot.png)

## Installation
See releases for [pre-built](https://github.com/rramiachraf/chorus/releases) binaries.
### Building from source:
#### Linux:
Dependencies:

- Go v1.18+
- Node.js v16+
- GCC
 
```bash
[~] $ git clone https://github.com/rramiachraf/chorus
[~] $ cd chorus
[chorus] $ make build-linux
```

#### Windows:
Dependencies:

- Go v1.18+
- Node.js v16+
- Mingw-w64
 
```
[~] $ git clone https://github.com/rramiachraf/chorus
[~] $ cd chorus
[chorus] $ make build-windows
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
[chorus] $ chorus MUSIC_DIR
```
**Note**: You can specify more directories by seperating them using a comma.
e.g. `chorus "MUSIC_DIR,MUSIC_DIR2"`.
