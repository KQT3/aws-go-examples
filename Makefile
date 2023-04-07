# Define the name of the executable file
APP_NAME=main

# Set the Go compiler to use
GO_COMPILER=go

# Set the source code directory
SRC_DIR=./cmd/server

# Set the output directory
OUT_DIR=./bin

# Define the build command
build:
	@echo "Building $(APP_NAME)..."
	$(GO_COMPILER) build -o $(OUT_DIR)/$(APP_NAME) $(SRC_DIR)

# Define the clean command
clean:
	@echo "Cleaning up..."
	rm -f $(OUT_DIR)/$(APP_NAME)

# Define the test command
test:
	@echo "Running tests..."
	$(GO_COMPILER) test -v ./...

# Define the default target
.PHONY: default
default: build

