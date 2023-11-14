build:
	go build -o bin/api 

run: build
	./bin/api --listenAddr :3000



test:
	go test -v ./...
