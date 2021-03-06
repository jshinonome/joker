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

var qConnPool = geek.NewConnPool()

var (
	basePort    = flag.Int("basePort", 1800, "The base q process port")
	gRPCPort    = flag.Int("gRPCPort", 1897, "The gRPC server port")
	qGinPort    = flag.Int("qGinPort", 1898, "The q gin server port")
	qEnginePort = flag.Int("qEnginePort", 1899, "The q engine port")
)

type trade struct {
	Time  time.Time `json:"time" k:"time"`
	Sym   string    `json:"sym" k:"sym"`
	Price float64   `json:"price" k:"price"`
	Qty   int64     `json:"qty" k:"qty"`
}

func main() {
	flag.Parse()
	initQConnPool()
	qEngine := geek.Engine{
		Port: *qEnginePort,
		Auth: func(u, p string) error { return nil },
		Pool: qConnPool,
	}
	qConnPool.Serving()
	log.Printf("geek engine listening at %v", qEngine.Port)

	// `::1899 `getTrade`a
	go qEngine.Run()
	r := gin.Default()
	r.GET("/trade/:sym", getTradeBySym)
	// curl http://localhost:1898/trade/a
	go r.Run(fmt.Sprintf(":%d", *qGinPort))
	// python ./pyclient/main.py
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *gRPCPort))
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
		qExec := exec.Command("q", "q/qprocess.q", "-p", strconv.Itoa(*basePort+i))
		qExec.SysProcAttr = &syscall.SysProcAttr{
			Pdeathsig: syscall.SIGTERM,
		}
		qExec.Start()
		time.Sleep(1 * time.Second)
		q := geek.QProcess{Port: *basePort + i}
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

type dataServer struct {
	api.UnimplementedDataServiceServer
}

func (s *dataServer) GetTrade(ctx context.Context, in *api.TradeRequest) (*api.TradeResponse, error) {
	log.Println("Got a gRPC message")
	start := time.Now()
	defer func() {
		log.Printf("[GRPC] %v", time.Since(start))
	}()
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
