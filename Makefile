OUT=chorus
VERSION=`git describe --tags`

init:
	cd view && yarn install && VITE_VERSION=$(VERSION) yarn build

build: build-linux

build-linux: init
	CGO_ENABLED=1 CC=gcc GOOS=linux GOARCH=amd64 go build -o dist/$(OUT)-$(VERSION)-linux_amd64 -ldflags="-X 'main.Version=$(VERSION)' -s -w" -tags sqlite_omit_load_extension

build-windows: init
	CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc GOOS=windows GOARCH=amd64 go build -o dist/$(OUT)-$(VERSION)-windows_amd64.exe -ldflags="-X 'main.Version=$(VERSION)' -s -w" -tags sqlite_omit_load_extension

build-all:
	rm -rf dist
	$(MAKE) build-linux
	$(MAKE) build-windows
	cd dist && sha256sum * > checksums.txt

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
