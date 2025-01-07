FROM arm64v8/golang:1.23.2-alpine3.20 AS builder

WORKDIR /app

COPY . /app

RUN go mod tidy && go build -o wish-bot

FROM alpine:latest

WORKDIR /app

RUN  apk --update add \
        ca-certificates \
        && \
        update-ca-certificates

COPY --from=builder /app/wish-bot /app/wish-bot

COPY --from=builder /app/config/config.yaml /app/config/config.yaml

ENTRYPOINT ["/app/wish-bot"]