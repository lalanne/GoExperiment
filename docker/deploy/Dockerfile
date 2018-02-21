FROM golang:1.10

WORKDIR /go/src/app

COPY server.go .

RUN go get -d -v github.com/go-sql-driver/mysql
RUN go get -d -v github.com/gorilla/mux

EXPOSE 8000

CMD go run server.go



