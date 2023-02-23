package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jshinonome/geek"
)

var q = geek.QProcess{Port: 1800}

func main() {
	q.Dial()
	r := gin.Default()
	r.GET("/hosts", getHosts)
	r.Run(":8080")
}

func getHosts(c *gin.Context) {
	h := make([]host, 0)
	if q.IsConnected() {
		q.Sync(&h, "hosts")
	}
	c.JSON(200, h)
}

type host struct {
	Host  string `json:"host" k:"host"`
	Port  int64  `json:"port" k:"port"`
	Label string `json:"label" k:"label"`
}
