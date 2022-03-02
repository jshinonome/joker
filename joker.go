/**
 * Copyright (c) 2022 Jo Shinonome
 *
 * This software is released under the MIT License.
 * https://opensource.org/licenses/MIT
 */
package main

import (
	"log"
	"net/http"
	"os/exec"
	"strconv"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jshinonome/geek"
)

var qConnPool geek.QConnPool

const qServerPort = 1899
const basePort = 1800

func main() {
	initQConnPool()
	qEngine := geek.QEngine{
		Port: qServerPort,
		Auth: func(u, p string) error { return nil },
		Pool: &qConnPool,
	}
	qConnPool.Serving()
	go qEngine.Run()
	r := gin.Default()
	// curl http://localhost:8080/trade/a
	r.GET("/trade/:sym", getTradeBySym)
	r.Run()
}

func initQConnPool() error {
	for i := 0; i < 2; i++ {
		qExec := exec.Command("q", "asset/q/qprocess.q", "-p", strconv.Itoa(basePort+i))
		qExec.SysProcAttr = &syscall.SysProcAttr{
			Pdeathsig: syscall.SIGTERM,
		}
		qExec.Start()
		time.Sleep(1 * time.Second)
		q := geek.QProcess{Port: basePort + i}
		err := q.Dial()
		if err != nil {
			return err
		}
		qConnPool.Put(&q)
	}
	return nil
}

func getTradeBySym(c *gin.Context) {
	sym := c.Param("sym")
	f := struct {
		Api string
		Sym string
	}{
		"getTrade", sym,
	}
	r := make([]trade, 0)
	err := qConnPool.Sync(&r, f)
	if err != nil {
		log.Println(err)
	}
	c.IndentedJSON(http.StatusOK, r)
}

type trade struct {
	Time  time.Time `json:"time" k:"time"`
	Sym   string    `json:"sym" k:"sym"`
	Price float64   `json:"price" k:"price"`
	Qty   int64     `json:"qty" k:"qty"`
}
