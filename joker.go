/**
 * Copyright (c) 2022 Jo Shinonome
 *
 * This software is released under the MIT License.
 * https://opensource.org/licenses/MIT
 */
package main

import (
	"log"
	"math/rand"
	"net/http"
	"os/exec"
	"strconv"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jshinonome/geek"
)

var qConnPool = make([]geek.QProcess, 2)

const basePort = 1800

func main() {
	initQConnPool()
	r := gin.Default()
	// curl http://localhost:8080/trade/a
	r.GET("/trade/:sym", getTradeBySym)
	r.Run()
}

func initQConnPool() error {
	for i := 0; i < len(qConnPool); i++ {
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
		qConnPool[i] = q
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
	i := rand.Intn(2)
	r := make([]trade, 0)
	err := qConnPool[i].Sync(&r, f)
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
