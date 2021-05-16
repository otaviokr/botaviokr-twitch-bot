FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git ca-certificates

WORKDIR /go/src/otaviokr/botaviokr-twitch-bot
COPY . .

RUN go get -d -v &&  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/botaviokr

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/bin/botaviokr /botaviokr

VOLUME [ "/config" ]
VOLUME [ "/logs" ]

ENTRYPOINT [ "/botaviokr" ]
