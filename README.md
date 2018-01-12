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
docker run --name MARIADB -p 3306:3306 -e MYSQL_ROOT_PASSWORD=pass -d mariadb:10.3.2
```

```
docker exec -it MARIADB bash
```

inside the container

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

 * TODO: timers for queries to DB, this should be the next step before SOAP.
 * TODO: investigate https://github.com/fiorix/wsdl2go to create SOAP client and server
```
wsdl2go < test1.wsdl > s.go
```

 * TODO: also check https://github.com/hooklift/gowsdl
