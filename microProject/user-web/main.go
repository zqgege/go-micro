package main

import (
        "fmt"
        "github.com/micro/cli"
        "github.com/micro/go-log"
        "github.com/micro/go-micro/registry"
        "github.com/micro/go-micro/registry/consul"
        "time"
        "microProject/basic"
        "microProject/basic/config"

        "github.com/micro/go-web"
        "microProject/user-web/handler"
)

func main() {
        basic.Init()

        //注册consul
        mireg := consul.NewRegistry(registryOptions)
        // 创建新服务
        service := web.NewService(
                // 后面两个web，第一个是指是web类型的服务，第二个是服务自身的名字
                web.Name("mu.micro.book.web.user"),
                web.Version("latest"),
                web.Registry(mireg),
                web.Address(":8088"),
        )
        // 初始化服务
        if err := service.Init(
                web.Action(
                        func(c *cli.Context) {
                        // 初始化handler
                        handler.Init()
                }),
        ); err != nil {
                log.Fatal(err)
        }

        // 注册登录接口
        service.HandleFunc("/user/login", handler.Login)
        service.HandleFunc("/user/logout", handler.Logout)

        // 运行服务
        if err := service.Run(); err != nil {
                log.Fatal(err)
        }
}

func registryOptions(ops *registry.Options) {
        consulCfg := config.GetConsulConfig()
        ops.Timeout = time.Second * 5
        ops.Addrs = []string{fmt.Sprintf("%s:%d", consulCfg.GetHost(), consulCfg.GetPort())}
}
