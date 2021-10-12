package main

import (
  "net/http"
  "github.com/gin-gonic/gin"

  "munvei/entry-exit-management/module"
)

func main () {
  r := gin.Default()

  r.GET("/", func(c *gin.Context) {
    res := module.DBConnect()
    dbStatus := "Failed"
    if res != nil {
      dbStatus = "Successful"
    }
    c.JSON(http.StatusOK, gin.H{
      "Status": "OK",
      "DB": dbStatus,
    })
  })

  r.Run() // listen and serve on 0.0.0.0:8080 <- default
}
