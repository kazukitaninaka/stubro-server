FROM golang:1.17-alpine as dev

RUN apk update && \
	apk add curl && \
	curl -L https://github.com/golang-migrate/migrate/releases/download/v4.11.0/migrate.linux-amd64.tar.gz | tar xvz && \
	mv ./migrate.linux-amd64 /usr/bin/migrate

RUN mkdir /app
WORKDIR /app

ENV GO111MODULE="on"
COPY go.mod go.sum /app/
RUN go mod download

ADD . /app
RUN go get github.com/pilu/fresh

FROM golang:1.17-alpine as prod

RUN apk update && \
	apk add curl && \
	curl -L https://github.com/golang-migrate/migrate/releases/download/v4.11.0/migrate.linux-amd64.tar.gz | tar xvz && \
	mv ./migrate.linux-amd64 /usr/bin/migrate

RUN mkdir /app
WORKDIR /app

ENV GO111MODULE="on"
COPY go.mod go.sum /app/
RUN go mod download

ADD . /app
RUN go build -o main .
CMD ["/app/main"]