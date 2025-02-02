FROM golang:1.23-alpine AS builder


WORKDIR /chat

COPY go.mod go.sum ./
RUN go mod download


COPY . .

RUN go build -o main

FROM alpine:latest


WORKDIR /root/

COPY --from=builder /chat/main .
COPY --from=builder /chat/template /template
COPY --from=builder /chat/static /static

EXPOSE 80


CMD ["./main"]
