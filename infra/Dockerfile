FROM golang:1.18-alpine as builder

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o ports ./cmd/ports.go

FROM alpine:latest

COPY --from=builder /app/ports .
CMD ["./ports"]
