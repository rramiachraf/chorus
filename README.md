# Chorus [![badge](https://github.com/rramiachraf/chorus/actions/workflows/main.yml/badge.svg)](https://github.com/rramiachraf/chorus/actions/workflows/main.yml)
Enjoy your music within the browser realm.

![Screenshot](https://raw.githubusercontent.com/rramiachraf/chorus/main/screenshot.png)

## Installation
Only Linux is currently supported.
### Building from source:
Dependencies:

- Go v1.18+
- Node.js v16+
- GCC
 
```bash
[~] $ git clone https://github.com/rramiachraf/chorus
[~] $ cd chorus
[chorus] $ make build
[chorus] $ sudo make install
```

## Usage
```bash
[chorus] $ chorus MUSIC_DIR
```
You can specify more directories by seperating them using a comma.
e.g. `./chorus "MUSIC_DIR,MUSIC_DIR2"`.

