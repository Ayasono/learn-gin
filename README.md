# learn-gin

## intro

### foundation of gin

#### create RESTful API

```go
server := gin.Default() // create a server

server.GET("{relative path}", handler) // get method
server.POST("{relative path}", handler) // post method
server.PUT("{relative path}", handler) // put method
server.DELETE("{relative path}", handler) // delete method
```

#### handler

```go
// handler, a function that handle the request
    func handler(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
           "message": "hello world",
    })
}
```

#### create router group

```go
    // access the router group by the relative path, e.g. /api/hello
    api := server.Group("/api") // create a router group
    api.GET("/hello", handler) // create a route in the group
```

#### middleware

```go
// middleware, a function that handle the request before the handler
func middleware(c *gin.Context) {
    // do something
    c.Next() // call the next middleware or handler
    c.abort() // abort the request
}

```
