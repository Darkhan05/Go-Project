FROM golang:1.24.2-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o main main.go

FROM alpine:latest

COPY --from=builder /app/main /root/

CMD ["/root/main"]
