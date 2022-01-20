# Go/Gin with MySQL template for Crafting Sandbox

This is a [Go](https://golang.org/)/[Gin](https://github.com/gin-gonic/gin) with MySQL template, configured for quick development setup in [Crafting Sandbox](https://docs.sandboxes.cloud/docs).

## Specifications

[`main.go`](main.go) specifies the port this package runs on, which matches the port set in App configuration:
```go
func main() {
	// ...
    
	ctx.Port = "3000"

	// ...
}
```

This template exposes a single `/ping` route:
```go
var routes = Routes{
	Route{http.MethodGet, "/ping", Ping},
}
```

This path accepts a query string, and responds with the query string and current time. For example:
```bash
$ curl --request GET 'localhost:3000/ping?ping=hello'
{"ping":"hello","received_at":"XXXX-XX-XX XX:XX:XX.XXXXXXXXX +0000 UTC"}
```

## App Definition

The following [App Definition](https://docs.sandboxes.cloud/docs/app-definition) was used to create this template:

```yaml
endpoints:
- name: api
  http:
    routes:
    - pathPrefix: "/"
      backend:
        target: go-gin
        port: api
    authProxy:
      disabled: true
workspaces:
- name: go-gin
  description: Template backend using Go/Gin
  ports:
  - name: api
    port: 3000
    protocol: HTTP/TCP
  checkouts:
  - path: backend
    repo:
      git: https://github.com/crafting-dev/template-go-gin
  packages:
  - name: golang
    version: 1.17.2
dependencies:
- name: mysql
  serviceType: mysql
  version: '8'
  properties:
    database: superhero
    password: batman
    username: brucewayne
```
