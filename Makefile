build:
	@go build -o out/bot ./cmd/main.go

run: build
	./out/bot