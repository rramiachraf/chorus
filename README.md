# Chorus
Enjoy your music within the browser realm.

## Installation
Go 1.18+ and Node.js 12+ are required.
```bash
[~] $ git clone https://github.com/rramiachraf/chorus
[~] $ cd chorus
[chorus] $ go build
[chorus] $ cd view
[chorus/view] $ yarn install
[chorus/view] $ yarn build
```

## Usage
```bash
[chorus] $ ./chorus MUSIC_DIR
```
You can specify more directories by seperating them using a comma.
e.g. `./chorus "MUSIC_DIR,MUSIC_DIR2"`.

## Notes
* You have a frontend where you will interact with your music and a backend which supplies data to your frontend, thus they need a way to connect with each other. By default the backend listens on port `3000`, on the other hand the frontend is a static app which resides on `view/public`, So you're going to need a web server to make the frontend able to talk with the backend. there's an nginx config file in the repo that you can use as a starting point or as reference for other web servers. 
