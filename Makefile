build:
	@go build -o out/bot ./cmd/main.go

run: build
	./out/bot serve

wh:
	curl -X POST "https://api.telegram.org/bot7186636734:AAGdMHLz619R07utv4t31J5shQDC1khIgbY/setWebhook" -d "url=https://ecombot-eight.vercel.app/api"