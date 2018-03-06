FROM golang:1.10

WORKDIR /go/src/app

COPY . .

RUN go get -d -v github.com/go-sql-driver/mysql
RUN go get -d -v github.com/gorilla/mux
RUN pwd
RUN ls -l
RUN ls -l web_service/
RUN ls -l web_service/src/
RUN export GOPATH=/go:/go/src/app/web_service && go build web_service/src/main.go

ENTRYPOINT ./main 

EXPOSE 8000


