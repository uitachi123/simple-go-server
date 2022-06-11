.PHONY: all test clean build run test
all: clean test build run

clean:
	rm -rf ./simple-go-server
build:
	go build -a -ldflags 'main.buildTime=$(date)' .
run:
	./simple-go-server --logging "DEBUG" --port "8080"
test:
	go test ./... -test.v
