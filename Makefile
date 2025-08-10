# Define default GOOS and GOARCH
GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)

# Specify the Go build command
BUILD_CMD = go build

# Default target to build for the current platform
all: build

# Build for the host platform
build:
	$(BUILD_CMD) -o bin/$(GOOS)_$(GOARCH)/temperature_sensor

# Build for Linux (64-bit)
build-linux:
	GOOS=linux GOARCH=amd64 $(BUILD_CMD) -o bin/linux_temperature_sensor

# Build for macOS (64-bit)
build-mac:
	GOOS=darwin GOARCH=amd64 $(BUILD_CMD) -o bin/mac_temperature_sensor

# Build for Windows (64-bit)
build-windows:
	GOOS=windows GOARCH=amd64 $(BUILD_CMD) -o bin/temperature_sensor.exe

# Generate the build for all platforms
build-all: build-linux build-mac build-windows
