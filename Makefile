OUT=chorus

all: dev

build:
	go build -o chorus
	cd view && yarn install && yarn build

dev:
	if ! command -v air &> /dev/null; then go install github.com/cosmtrek/air@v1.42.0; fi
	air &
	cd view && if ! test -d node_modules; then yarn install; fi && yarn dev
