package main

import (
	"fmt"
	"github.com/micro/cli"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"microProject/auth/handler"
	"microProject/auth/model"
	"microProject/basic"
	"microProject/basic/config"
	"time"

	s "microProject/auth/proto/auth"
)

func main() {
	basic.Init()

	miReg := consul.NewRegistry(registryOptions)
	// New Service
	service := micro.NewService(
		micro.Name("mu.micro.book.srv.auth"),
		micro.Registry(miReg),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init(
		micro.Action(func(c *cli.Context) {
			model.Init()
			handler.Init()
		}),
	)

	//注册服务
	s.RegisterServiceHandler(service.Server(), new(handler.Service))
	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func registryOptions(ops *registry.Options) {
	consulCfg := config.GetConsulConfig()
	ops.Timeout = time.Second * 5
	ops.Addrs = []string{fmt.Sprintf("%s:%d", consulCfg.GetHost(), consulCfg.GetPort())}
}
