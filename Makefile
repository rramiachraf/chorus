OUT=chorus

all: dev

build:
	go build -o chorus
	cd view && yarn install && yarn build

dev:
	if ! command -v air &> /dev/null; then go install github.com/cosmtrek/air@v1.42.0; fi
	if ! command -v overmind &> /dev/null; then go install github.com/DarthSim/overmind@v2; fi
	overmind start
