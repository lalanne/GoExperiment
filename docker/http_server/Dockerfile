FROM golang:1.10

WORKDIR /go/src/app

COPY http_server/http.go .

RUN go get -d -v github.com/gorilla/mux
RUN go build http.go

ENTRYPOINT ./http 

EXPOSE 8080


