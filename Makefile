BINARY_NAME=blogrender
CMD_PATH=./cmd/blogrender
BUILD_DIR=bin

build:
	go build -mod=vendor -o $(BUILD_DIR)/$(BINARY_NAME) $(CMD_PATH)

# Linux build
build-linux:
	GOOS=linux GOARCH=amd64 go build -mod=vendor -o $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64 $(CMD_PATH)

# Windows build  
build-windows:
	GOOS=windows GOARCH=amd64 go build -mod=vendor -o $(BUILD_DIR)/$(BINARY_NAME)-windows-amd64.exe $(CMD_PATH)

# macOS build (Intel)
build-macos:
	GOOS=darwin GOARCH=amd64 go build -mod=vendor -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-amd64 $(CMD_PATH)

# macOS build (Apple Silicon M1/M2)
build-macos-arm:
	GOOS=darwin GOARCH=arm64 go build -mod=vendor -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-arm64 $(CMD_PATH)

release: build-linux build-windows build-macos build-macos-arm
	@echo "Built binaries for all platforms in $(BUILD_DIR)/"
	@ls -la $(BUILD_DIR)/

install:
	go install -mod=vendor $(CMD_PATH)

run:
	go run -mod=vendor $(CMD_PATH)

clean:
	rm -rf $(BUILD_DIR)

test:
	go test ./...

benchmark:
	go test -bench=. ./... -count 5 -benchtime=10s -benchmem