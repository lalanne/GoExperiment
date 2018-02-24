# GoExperiment

The idea is to create a service that connect to a Database and to several external 
web services via **SOAP** and plaing **HTTP**, exploring concurrency and networking
in golang.
```
http.ListenAndServe(":8000", router)
``` 
This creates an **http server** listening on port **8080**, for what I found every 
handler function it is executed in its own **goroutine**.

Also the service uses **gorilla** web toolkit as the service **multiplexer**
(it is compatible with the standard http handler API)

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
To implement timers for queries to the data base one way to do it, is using
```
readTimeout
```
This is a timer to query the database for just for reading queries, it is configured 
for db Connection, **NOT** by query.

Because we wanted to configure timers by query, we used the **context** package, the
DB package implements the interface ```QueryContext``` which previously it is
configured by this function ```context.WithTimeout```

## Using docker to start the system

Build the container for the database:
```
docker build --no-cache -t clalanne/goexperiment_db -f ./docker/db/Dockerfile .
```
Build the container for the service:
```
docker build --no-cache -t clalanne/goexperiment -f ./docker/deploy/Dockerfile .
```
Push db container to registry:
```
docker push clalanne/goexperiment_db:latest
```
Push service container to registry:
```
docker push clalanne/goexperiment:latest
```
Run db container:
```
docker run -it --name GOEXPERIMENT_DB -p 3306:3306 -e MYSQL_ROOT_PASSWORD=pass -d clalanne/goexperiment_db:latest
```
Run service container:
```
docker run -it --name GOEXPERIMENT -p 8000:8000 clalanne/goexperiment:latest
```
Enter to the containers
```
docker exec -it GOEXPERIMENT_DB bash
```
```
docker exec -it GOEXPERIMENT bash
```
Once inside db container to enter to the already created database:
```
mysql -u root -ppass
```

## docker-compose
To start services
```
docker-compose up -d
```
To enter to ws container
```
docker exec -it goexperiment_ws_1 bash
```
To test ws, from host
```
curl "http://0.0.0.0:8000/purchase?a=5&b=9"
```
To stop services
```
docker-compose down
```
## Example command for request http to service:
```
curl "http://127.0.0.1:8000/purchase?a=5&b=9"
```

## Notes
Seems that there is no reason to use asynchronous querying to the database, because
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

