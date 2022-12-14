build:
	go build -o bin/skeleton cmd/main.go

run: build
	export APP_ENV=dev && ./bin/skeleton