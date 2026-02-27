.PHONY: build install clean

# Build the binary
build:
	go build -o git-pick main.go

# Install globally (requires sudo on Linux, admin on macOS)
install: build
	mv git-pick /usr/local/bin/git-pick
	chmod +x /usr/local/bin/git-pick

# Clean build artifacts
clean:
	rm -f git-pick

# Uninstall
uninstall:
	rm -f /usr/local/bin/git-pick
