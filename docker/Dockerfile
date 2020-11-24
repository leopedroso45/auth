# Imagem base
FROM golang:1.14

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

#RUN chmod +x wait-for-it.sh

ENTRYPOINT go run .