package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/hosts", getHosts)
	r.Run(":8080")
}

func getHosts(c *gin.Context) {
	h := []host{
		{Host: "localhost", Port: 2000, Label: "tca"},
		{Host: "127.0.0.1", Port: 2001, Label: "md"},
		{Host: "localhost", Port: 2002, Label: "vc"},
		{Host: "localhost", Port: 2003, Label: "config"},
		{Host: "localhost", Port: 2004, Label: "calendar"},
	}
	c.JSON(200, h)
}

type host struct {
	Host  string `json:"host"`
	Port  int64  `json:"port"`
	Label string `json:"label"`
}
