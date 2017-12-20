# GoExperiment

The idea is to create a service that connect to a Database and to an external service
via SOAP, exploring concurrency and networking in golang.

```http.ListenAndServe(":8000", router)``` This creates an http server listening on port
8080, for what I found every handler function it is executed in its own goroutine.

Also the service uses gorilla web toolkit as the service multiplexer(it is compatible
with the standard http handler API)

It is necessary to use a database, the easiest way is to use a docker container:

```
docker run --name MARIADB -e MYSQL_ROOT_PASSWORD=pass -d mariadb:10.3.2
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
