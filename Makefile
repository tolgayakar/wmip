.PHONY: build install clean

# Define the output binary path
BINARY_NAME = wmip
OUTPUT_DIR = bin

# Build the application
build:
	@mkdir -p $(OUTPUT_DIR)
	go build -o $(OUTPUT_DIR)/$(BINARY_NAME) cmd/main.go

# Install the application
install: build
	cp $(OUTPUT_DIR)/$(BINARY_NAME) /usr/local/bin/

# Clean up build artifacts
clean:
	rm -f $(OUTPUT_DIR)/$(BINARY_NAME)
