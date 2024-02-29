FROM golang:1.22.0
WORKDIR /build
COPY go.mod .
RUN go mod download
COPY . .
RUN go build
CMD [ "make", "start_prod" ]