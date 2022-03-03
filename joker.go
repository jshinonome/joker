/**
 * Copyright (c) 2022 Jo Shinonome
 *
 * This software is released under the MIT License.
 * https://opensource.org/licenses/MIT
 */
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os/exec"
	"strconv"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jshinonome/geek"
	"github.com/jshinonome/joker/api"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var qConnPool geek.QConnPool

const qEnginePort = 1899
const qGinPort = 1898
const basePort = 1800

var (
	port = flag.Int("port", 1897, "The gRPC server port")
)

type dataServer struct {
	api.UnimplementedDataServiceServer
}

// python ./pyclient/main.py
func (s *dataServer) GetTrade(ctx context.Context, in *api.TradeRequest) (*api.TradeResponse, error) {
	log.Println("Got a gRPC message")
	sym := in.GetSym()
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
	trades := make([]*api.Trade, len(r))
	for i, t := range r {
		trades[i] = &api.Trade{
			Time:  timestamppb.New(t.Time),
			Sym:   t.Sym,
			Price: t.Price,
			Qty:   t.Qty,
		}
	}
	return &api.TradeResponse{Trades: trades}, nil
}

func main() {
	initQConnPool()
	qEngine := geek.QEngine{
		Port: qEnginePort,
		Auth: func(u, p string) error { return nil },
		Pool: &qConnPool,
	}
	qConnPool.Serving()
	go qEngine.Run()
	r := gin.Default()
	// curl http://localhost:1898/trade/a
	r.GET("/trade/:sym", getTradeBySym)
	go r.Run(fmt.Sprintf(":%d", qGinPort))

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	api.RegisterDataServiceServer(s, &dataServer{})
	log.Printf("grpc server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
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
