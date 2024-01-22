build:
	go build -o bin/golaundry

run: build
	./bin/golaundry

test:
	go test -v ./...