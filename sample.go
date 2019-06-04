package main

import (
  "github.com/gin-gonic/gin"
  "net/http"
)

func main() {
  router := gin.Default()
  router.LoadHTMLGlob("templates/*.tmpl")

  router.GET("/", func(c *gin.Context) {
    c.HTML(http.StatusOK, "index.tmpl", gin.H{ "as": "2222" })
  })
  router.GET("/lg", ping)
  router.Run(":8080")
}

func ping(c *gin.Context) {
  c.HTML(http.StatusOK, "index.tmpl", gin.H{
    "as": "2222",
    "result": "hoge",
  })
}
