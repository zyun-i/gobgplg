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
    c.HTML(http.StatusOK, "index.tmpl", gin.H{ "as": "64565", "checked": "0"})
  })
  router.GET("/lg", query)
  router.Run(":8080")
}

func query(c *gin.Context) {

  query := c.Query("query")
  arg := string(c.Query("arg"))
  var cmd_name string
  var cmd_arg []string

  switch query {
  case "0":
    cmd_name = "bgpctl"
    cmd_arg = []string{"show"}
  case "1":
    cmd_name = "bgpctl"
    cmd_arg = []string{"show", "neighbor"}
  case "2":
    cmd_name = "bgpctl"
    cmd_arg = []string{"show", "nexthop"}
  case "3":
    cmd_name = "bgpctl"
    cmd_arg = []string{"show", "ip", "bgp"}
  case "4":
    cmd_name = "bgpctl"
    cmd_arg = []string{"show", "ip", "bgp", arg}
  case "5":
    cmd_name = "bgpctl"
    cmd_arg = []string{"show", "ip", "bgp", "as", arg}
  case "6":
    cmd_name = "ping"
    cmd_arg = []string{"-c4", "-i0.5", arg}
  case "7":
    cmd_name = "traceroute"
    cmd_arg = []string{"-w2", "-q2", "-m25", arg}
  case "8":
    cmd_name = "ping6"
    cmd_arg = []string{"-c4", "-i0.5", arg}
  case "9":
    cmd_name = "traceroute6"
    cmd_arg = []string{"-w2", "-q2", "-m25", arg}
  }

  out, err := exec.Command(cmd_name, cmd_arg...).Output()
  if err != nil {
    fmt.Printf("Failed Foo: %s",err)
  }

  c.HTML(http.StatusOK, "index.tmpl", gin.H{
    "as": "64565",
    "result": string(out),
    "checked": string(query),
    "arg": string(arg),
  })
}
