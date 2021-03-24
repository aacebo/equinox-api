build:
	go build -o bin/main src/main.go

clean:
	rm -rf ./bin

run:
	GO_ENV=local go run src/main.go
