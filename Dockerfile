FROM golang:1.22-alpine3.19 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o minecraft-discord-bot .

FROM alpine:3.19 AS certs

RUN apk add --no-cache ca-certificates

FROM scratch

# api.mcsrvstat.us use a CA not trusted by default in the scratch image
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /app/minecraft-discord-bot /app/minecraft-discord-bot

ENTRYPOINT ["/app/minecraft-discord-bot"]
