FROM golang:1.22-alpine3.20 AS builder

COPY . github.com/DenisCom3/m-chat-server
WORKDIR github.com/DenisCom3/m-chat-server

RUN go mod download
RUN go build -o ./bin/m-chat-server cmd/main.go

FROM alpine:3.20

WORKDIR /root/
COPY --from=builder /go/github.com/DenisCom3/m-chat-server/bin .
COPY --from=builder /go/github.com/DenisCom3/m-chat-server/.env .

CMD ["./m-chat-server", "--config-path", ".env" ]