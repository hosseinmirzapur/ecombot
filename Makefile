build:
	@go build -o out/bot

run: build
	./out/bot