FROM golang:1.17-alpine

RUN mkdir /app

RUN go get \
    firebase.google.com/go/v4 \
    github.com/gorilla/mux \
    google.golang.org/api \
    gorm.io/driver/mysql \
    gorm.io/gorm

ADD . /app
WORKDIR /app
RUN go build -o main .
CMD ["/app/main"]