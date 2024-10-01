BINARY_NAME = brain-game
MAIN_FILE = ./internal/main.go

build:
	go build -o internal/$(BINARY_NAME) $(MAIN_FILE)

run: build
	./internal/$(BINARY_NAME)

clean:
	rm ./internal/$(BINARY_NAME)

lint:
	golangci-lint run

test:
	go test ./tests/... -v

install:
	go mod tidy