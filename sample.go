package main

import (
  "fmt"
  "github.com/gin-gonic/gin"
  "net/http"
  "os/exec"
)

func main() {
  router := gin.Default()
  router.LoadHTMLGlob("templates/*.tmpl")

  router.GET("/", func(c *gin.Context) {
    c.HTML(http.StatusOK, "index.tmpl", gin.H{ "as": "64565" })
  })
  router.GET("/lg", ping)
  router.Run(":8080")
}

func ping(c *gin.Context) {

  query := c.Query("query")
  var cmd string

  switch query {
  case "0":
    cmd = "summary"
  case "1":
    cmd = "ip bgp"
  }

  out, err := exec.Command("bgpctl", "show", cmd ).Output()
  if err != nil {
	  fmt.Printf("Failed Foo: %s",err)
  }

  c.HTML(http.StatusOK, "index.tmpl", gin.H{
    "as": "64565",
    "result": string(out),
  })
}
