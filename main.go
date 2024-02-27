package main

import (
  "encoding/json"
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
    b, _ := c.GetRawData()

    jsonData := make(map[string]interface{})

    _ = json.Unmarshal(b, &jsonData)

    c.JSON(http.StatusOK, gin.H{
      "msg":   "ok",
      "value": jsonData["test"],
    })
  })

  server.GET("/user/:name", func(c *gin.Context) {
    name := c.Param("name")

    for _, user := range users {
      if user.Name == name {
        c.JSON(http.StatusOK, gin.H{
          "user": user,
        })
        return
      }
    }

    c.JSON(http.StatusNotFound, gin.H{
      "message": "User not found",
    })
  })

  orderGroup := server.Group("/order")
  {
    orderGroup.GET("/list", func(c *gin.Context) {
      c.JSON(http.StatusOK, gin.H{
        "orders": "list",
      })

    })
    orderGroup.GET("/detail", func(c *gin.Context) {
      c.JSON(http.StatusOK, gin.H{
        "orders": "detail",
      })
    })
  }

  server.NoRoute(func(c *gin.Context) {
    c.JSON(http.StatusNotFound, gin.H{
      "message": "Not found",
    })
  })

  server.Run(":8080")
}
