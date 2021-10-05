# Go/Gin with MySQL template for Cloud Sandbox

This is a simple [Gin](https://github.com/gin-gonic/gin) framework template for developing applications in Cloud Sandbox.

## Specification

The following [App](https://crafting.readme.io/docs/app-spec) configuration was used in Cloud Sandbox to create this template:

```yaml
spec:
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
        root-password: batman
      service_type: mysql
      version: "8"
    name: mysql
```