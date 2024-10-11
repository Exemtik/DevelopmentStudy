BINARY_NAME = brain-game
MAIN_FILE = ./cmd/main.go

build:
	go build -o cmd/$(BINARY_NAME) $(MAIN_FILE)

run: build
	./cmd/$(BINARY_NAME)

clean:
	rm ./cmd/$(BINARY_NAME)

lint:
	golangci-lint run

test:
	go test ./tests/... -v

install:
	go mod tidy