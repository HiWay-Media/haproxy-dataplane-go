---
layout: default
title: Home
nav_order: 1
description: "Haproxy Dataplane Go lang Docs"
permalink: /
last_modified_date: 2022-11-07T21:54:08+0000
---

# Haproxy Dataplane Go lang 

A client library for interacting with the Haproxy Dataplane API V2 for listing, creating, updating and delete haproxy definitions like:

 - basic informations
 - backends
 - frontends
 -  backends and frontends rules
 -  servers
 - acls
 - http rules
 - filters
 and more.

### Example usage:

```go
// creating a client
client := haproxy.NewHaproxyClient("127.0.0.1", "user", "password", true)

// listing backends
backends, err := client.GetBackends()

//listing acls (eg: parent type, parent name)
servers, err := client.GetAcls("frontend", "foo")
```

for other informations refer to the HaProxy Dataplane V2 API spec.
