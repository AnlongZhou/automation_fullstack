
FROM golang:1.25.1 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main ./cmd/main.go

FROM ubuntu:22.04

WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/views ./views
COPY --from=builder /app/css ./css

ENV PORT=3000
EXPOSE 3000

CMD ["./main"]
