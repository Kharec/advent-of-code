GO = go
BIN = bin/aoc

.PHONY: build test clean

build:
	$(GO) build -o $(BIN) ./cmd/aoc

test:
	$(GO) test ./...

clean:
	rm -f $(BIN)
