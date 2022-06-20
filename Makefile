LEVEL?="INFO"

.PHONY: all test clean build api ui serve test
all: clean test build serve

clean:
	rm -rf ./simple-go-server
build:
	go build -a -ldflags 'main.buildTime=$(date)' .
api:
	./simple-go-server --logging $(LEVEL) --port "8080"
ui:
	npm run build && npm start
serve:
	make -j api ui
test:
	go test ./... -test.v
