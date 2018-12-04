# Version: 1.0.0

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build -v

all: clean build
build:
	# Update/Download libraries
	$(GOCMD) mod download
	# Build
	$(GOBUILD) -o rename main.go
clean:
	rm -rf rename
