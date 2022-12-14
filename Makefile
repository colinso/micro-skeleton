build:
	go build -o bin/skeleton cmd/main.go

run: build
	./bin/skeleton