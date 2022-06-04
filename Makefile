.PHONY: all test clean build run test
all: clean test build run

clean:
	rm -rf ./go-server
build:
	go build -a -ldflags 'main.buildTime=$(date)' .
run:
	./go-server --logging "DEBUG" --port "8080"
test:
	go test ./... -test.v
