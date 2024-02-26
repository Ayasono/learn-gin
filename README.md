# learn-gin

## intro

### foundation of gin

> use get method report a json

```go
func main() {
  server := gin.Default()

  server.GET("/user", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "msg": "ok",
    })
  })
  
  server.Run(":8080")
}
```

> use post method report a json

```go
func main() {
  server := gin.Default()

  server.POST("/json", func(c *gin.Context) {
    var jsonData map[string]interface{}
    
    // 使用c.BindJSON来解析请求体到一个map中
    if err := c.BindJSON(&jsonData); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
    }
    
    // 遍历map，并返回所有键值对
    c.JSON(http.StatusOK, jsonData)
  
  server.Run(":8080")
}
```
