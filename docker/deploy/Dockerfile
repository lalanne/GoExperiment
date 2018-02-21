FROM golang:1.10

WORKDIR /go/src/app

COPY server.go .

RUN go get -d -v github.com/go-sql-driver/mysql
RUN go get -d -v github.com/gorilla/mux
RUN go build server.go

ENTRYPOINT ./server

EXPOSE 8000


