OUT=chorus
VERSION=`git describe --tags`

all: dev

build:
	cd view && yarn install && VITE_VERSION=$(VERSION) yarn build
	CGO_ENABLED=1 CC=gcc GOOS=linux GOARCH=amd64 go build -v -o $(OUT) -ldflags="-X 'main.Version=$(VERSION)' -s -w" -tags sqlite_omit_load_extension

dev:
	if ! command -v air &> /dev/null; then go install github.com/cosmtrek/air@v1.42.0; fi
	if ! command -v overmind &> /dev/null; then go install github.com/DarthSim/overmind@v2; fi
	overmind start

test:
	go test ./... -v -cover

install:
	install -Dm755 $(OUT) -t /usr/bin
	install -Dm644 README.md -t /usr/share/doc/$(OUT)
	install -Dm644 LICENSE -t /usr/share/licenses/$(OUT)
