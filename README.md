# Go/Gin with MySQL template for Crafting Sandbox

This is a [Go](https://golang.org/)/[Gin](https://github.com/gin-gonic/gin) with MySQL template, configured for quick development setup in Crafting Sandbox.

## Specifications

[`main.go`](main.go) specifies the port this package runs on, which matches the port set in App configuration:
```go
func main() {
	// ...
    
	ctx.Port = "3000"

	// ...
}
```

All routes are grouped and versioned by number (for example: `baseUrl/v1` where `v1` is the api version). See Gin's [docs](https://github.com/gin-gonic/gin#grouping-routes) on grouping routes.

This template exposes a single `/ping` route:
```go
var routes = Routes{
	Route{http.MethodGet, "/ping", PingPong},
}
```

This path accepts a query string, and responds with the query string and current time. For example:
```bash
$ curl --request GET 'localhost:3000/v1/ping?ping=hello'
{"ping":"hello","current_time":"XXXX-XX-XX XX:XX:XX.XXXXXXXXX +0000 UTC"}
```



## App configuration

The following [App configuration](https://crafting.readme.io/docs/app-spec) was used to create this template:

```yaml
endpoints:
- http:
  routes:
  - backend:
      port: http
      target: go-gin
    path_prefix: /
name: api
services:
- description: Go/Gin template
name: go-gin
workspace:
  checkouts:
  - path: src/template-go-gin
    repo:
      git: https://github.com/crafting-dev/template-go-gin.git
  packages:
  - name: golang
    version: ~1.17
  ports:
  - name: http
    port: 3000
    protocol: HTTP/TCP
- managed_service:
  properties:
    database: superhero
    password: batman
    username: brucewayne
  service_type: mysql
  version: "8"
name: mysql
```
