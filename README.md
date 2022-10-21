# haproxy-dataplane-go

[![Go Report Card](https://goreportcard.com/badge/github.com/Allan-Nava/haproxy-dataplane-go)](https://goreportcard.com/report/github.com/Allan-Nava/haproxy-dataplane-go)
[![GoDoc](https://godoc.org/github.com/Allan-Nava/haproxy-dataplane-go?status.svg)](https://godoc.org/github.com/Allan-Nava/haproxy-dataplane-go)

### go client for HaProxy Dataplane API

------------

A client library for interacting with the Haproxy Dataplane API V2 for 
listing, creating, updating and delete haproxy definitions like:

 - basic informations
 - backends
 - frontends
 -  backends and frontends rules
 -  servers
 - acls
 - http rules
 - filters
 and more.

------------

### example usage:

```go
// creating a client
client, err := haproxy.NewClient("127.0.0.1", "user", "password", true)

// listing backends
backends, err := client.GetBackends()

//listing acls (eg: parent type, parent name)
servers, err := client.GetAcls("frontend", "foo")
```

for other informations refer to the HaProxy Dataplane V2 API spec.

## WORK IN PROGRESS

------------
#### license 
- MIT
