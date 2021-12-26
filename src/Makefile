.PHONY: run test build build-win build-mac

run:
	echo "Run go application"
	go fmt ./...
	clear && go run .

test:
	go test -v ./test

build:
	echo "Build go application for Linux OS..."
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./dist/otoklix-blog

build-win:
	echo "Build go application for Windows OS..."
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ./dist/otoklix-blog

build-mac:
	echo "Build go application for Mac OS..."
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ./dist/otoklix-blog

