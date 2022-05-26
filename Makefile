all: clean build run

clean:
	rm -rf ./go-server
build:
	go build -a -ldflags 'main.buildTime=$(date)' .
run:
	./go-server --logging "DEBUG"
