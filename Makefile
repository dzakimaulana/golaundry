BINARY_NAME=golaundry
BUILD_DIR=./bin
CMD_DIR=./cmd/golaundry

build:
	go build -o ${BUILD_DIR}/${BINARY_NAME} ${CMD_DIR}/main.go

run: build
	${BUILD_DIR}/${BINARY_NAME}

test:
	go test -v ./internal/...

clean:
	go clean
	del /q ${BUILD_DIR}/${BINARY_NAME}