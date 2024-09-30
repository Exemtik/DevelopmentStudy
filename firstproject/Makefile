BINARY_NAME = file-game
MAIN_FILE = bin/main.go

build:
	go build -o bin/$(BINARY_NAME) $(MAIN_FILE)

run: build
	./bin/$(BINARY_NAME)

clean:
	rm ./bin/$(BINARY_NAME)

install:
	go mod tidym