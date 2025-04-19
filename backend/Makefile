# Makefile for compiling and running the Go application

# Set the binary name
BINARY_NAME=gamify

# Build the Go binary
build:
	go build -o $(BINARY_NAME) ./cmd/gamify

# Run the Go application
run: build
	./$(BINARY_NAME)

# Clean up compiled binaries
clean:
	rm -f $(BINARY_NAME)

# Install the dependencies
deps:
	go mod tidy

# Run tests
test:
	go test ./...

# Build and run the application (shortcut)
all: build run
