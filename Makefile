build:
	@go build -o ./bin/spxce

run: build
	@./bin/spxce

test:
	@go test ./...