# Go/Gin with MySQL template for Cloud Sandbox

This is a Go [Gin](https://github.com/gin-gonic/gin) template, configured with MySQL, for quick development setup in Cloud Sandbox.

## Specifications

The following `App configuration` was used to create this template:

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
- description: A gin-gonic api template
  name: go-gin
  workspace:
    checkouts:
    - path: template-go-gin
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
      username: batman
      password: brucewayne
      database: app
    service_type: mysql
    version: "8"
  name: mysql
```