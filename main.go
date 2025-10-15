package main

import (
	"flag"

	"github.com/xucxapp/gin-gateway/internal/gateway"
)

var configFile = flag.String("f", "etc/gateway.yaml", "配置文件")

func main() {
	flag.Parse()

	server := &gateway.Server{}
	if err := server.LoadConfig(*configFile); err != nil {
		panic(err)
	}
	// defer server.Stop()
	server.Start()
}
