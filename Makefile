all: clean build run

clean:
	rm -rf ./go-server
build:
	go build -a .
run:
	./go-server
