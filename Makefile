.PHONY: build proto test run

build:
	go build -o bin/server cmd/server/main.go

proto:
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		proto/cache/v1/cache.proto

test:
	go test ./...

run:
	./bin/server
