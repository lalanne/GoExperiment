# GoExperiment

The idea is to create a service that connect to a Database and to an external service
via SOAP, exploring concurrency and networking in golang.

```http.ListenAndServe(":8000", router)``` This creates an http server listening on port
8080, for what I found every handler function it is executed in its own goroutine.

Also the service uses gorilla web toolkit as the service multiplexer(it is compatible
with the standard http handler API)

For the service to run some packages are needed:
 * go mysql driver:
 ```
 go get -u github.com/go-sql-driver/mysql
 ```
  * http multiplexer:
```
go get -u github.com/gorilla/mux
```
## Database

It is necessary to use a database, the easiest way is to use a docker container:

```
docker build --no-cache -t registry:5000/goexperiment_db -f ./docker/db/Dockerfile .
```
```
docker push registry:5000/goexperiment_db:latest
```
```
docker run -it --name GOEXPERIMENT_DB -p 3306:3306 -e MYSQL_ROOT_PASSWORD=pass -d registry:5000/goexperiment_db:latest
```
```
docker exec -it GOEXPERIMENT_DB bash
```
```
mysql -u root -ppass
```

## Example command for request http to service:
```
curl "http://127.0.0.1:8000/purchase?a=5&b=9"
```

## Notes
Seems that there is no reason to use asynchronous querying of the database, because
requests are handled by goroutines and this are switched when there is IO blocking 
involved, so there is no blocking when querying the database, it seems to me that this
approach is even better than state machine non-blocking approach because its easier
to reason about

# SOAP options
## wsdl2go
  * https://github.com/fiorix/wsdl2go 
To create SOAP client and server
```
wsdl2go < test1.wsdl > s.go
```
with ```wsdl2go``` test.wsdl does not work only test1.wsdl

## gowsdl
 * https://github.com/hooklift/gowsdl

With gowsdl both wsdl on the repo compile correctly and go code is generated in 
both cases

