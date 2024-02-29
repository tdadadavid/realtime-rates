# == Build Stage ==
FROM golang:1.22.0 as builder 

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

ENV GOOS=linux
ENV GOARCH=amd64

RUN go build -o main

# == Final Stage ==
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main /app/

EXPOSE 3000

CMD [ "./main" ] //TODO work on this.