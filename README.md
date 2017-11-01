# GoExperiment

The idea is to create a service that connect to a Database and to an external service
via SOAP, exploring concurrency and networking in golang.

```http.ListenAndServe(":8000", router)``` This creates an http server listening on port
8080, for what I found every handler function it is executed in its own goroutine.

Also the service uses gorilla web toolkit as the service multiplexer(it is compatible
with the standard http handler API)
