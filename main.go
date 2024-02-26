package main

import (
  "net/http"

  "github.com/gin-gonic/gin"
)

func main() {
  server := gin.Default()

  type userInfo struct {
    Name     string `json:"name"`
    Password string `json:"password"`
  }

  var users []userInfo

  server.GET("/user", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "users": users,
    })
  })

  server.POST("/user", func(c *gin.Context) {
    userName := c.Query("name")
    userPassword := c.Query("password")

    users = append(users, userInfo{Name: userName, Password: userPassword})

    c.JSON(http.StatusOK, gin.H{
      "message": "User added successfully",
      "users":   users,
    })
  })

  server.POST("/json", func(c *gin.Context) {
    var jsonData map[string]interface{}

    // 使用c.BindJSON来解析请求体到一个map中
    if err := c.BindJSON(&jsonData); err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      return
    }

    // 遍历map，并返回所有键值对
    c.JSON(http.StatusOK, jsonData)
  })

  server.Run(":8080")
}
