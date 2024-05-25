FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY . .
RUN GOPROXY=https://goproxy.cn,direct GO111MODULE=on go build -o /app/ecombot cmd/main.go
COPY . .
EXPOSE 8080
CMD [ "/app/ecombot", "serve" ]