# Imagem base
FROM golang

LABEL maintainer="leopedroso45 <leopedroso45@gmail.com>"

WORKDIR /app/src/auth

ENV GOPATH=/app

COPY . /app/src/auth

RUN go build main.go

ENTRYPOINT ["./main"]

EXPOSE 8080