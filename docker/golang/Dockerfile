ARG GO_VERSION=1.25

FROM golang:${GO_VERSION}-alpine as builder

WORKDIR /application

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o main ./cmd/api

FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /root/

COPY --from=builder /application/config ./config
COPY --from=builder /application/main .

CMD ["./main"]