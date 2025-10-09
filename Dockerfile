
FROM golang:1.25.1 AS base

RUN go install github.com/air-verse/air@latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

ENV PORT=3000
EXPOSE 3000

CMD ["air", "-c", ".air.toml"]
