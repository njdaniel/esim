.PHONEY: build run

build:
	@echo "Building..."
	env GOOS=linux GOARCH=amd64 go build --tags netgo --ldflags '-extldflags"lm -lstdc++ -static"'

run:
	go run main.go