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

WORKDIR /

COPY --from=builder /app/main /app/

# Set execute permission
RUN chmod +x /app/main

# Add any necessary dependencies (if needed)
RUN apk --no-cache add libc6-compat

EXPOSE 3000

CMD [ "./app/main" ]