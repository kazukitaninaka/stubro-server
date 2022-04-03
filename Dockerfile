FROM golang:1.17-alpine

RUN mkdir /app
WORKDIR /app

ENV GO111MODULE="on"
COPY go.mod go.sum /app/
RUN go mod download

ADD . /app
RUN go build -o main .
CMD ["/app/main"]