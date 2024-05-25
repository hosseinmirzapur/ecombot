FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY . .
RUN GOPROXY=https://goproxy.cn,direct GO111MODULE=on go build -o /app/ecombot cmd/main.go
COPY . .
EXPOSE 8090
CMD [ "/app/ecombot", "serve", "--http=0.0.0.0:8090", "--dir=/opt/pb_data" ]