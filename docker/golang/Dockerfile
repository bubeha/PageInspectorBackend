# Убираем дублирование ARG
ARG GO_VERSION=1.25

FROM golang:${GO_VERSION}-alpine AS builder

RUN apk add --no-cache shadow && useradd -u 1001 forge

WORKDIR /application

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o main ./cmd/api

FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app

COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /application/main ./
COPY --from=builder /application/config ./config
COPY --from=builder --chown=forge /application/certs/jwt-public.pem /etc/application/keys/jwt-public.pem
COPY --from=builder --chown=forge /application/certs/jwt-private.pem /etc/application/keys/jwt-private.pem

USER forge

CMD ["./main"]
