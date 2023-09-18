package main

import (
	"ToDoList/app/gateway/router"
	"ToDoList/app/gateway/rpc"
	"ToDoList/config"
	"fmt"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/web"
	"time"
)

func main() {
	//初始化配置文件等的读取
	config.Init()
	rpc.InitRPC()
	//注册一个etcd
	etcdReg := registry.NewRegistry(registry.Addrs(fmt.Sprintf("%s:%s", config.EtcdHost, config.EtcdPort)))
	//把当前的服务注册成一个微服务
	webService := web.NewService(
		web.Name("httpService"),
		web.Address(config.HttpServiceAddress), //服务的地址
		web.Registry(etcdReg),                  //etcd服务注册
		web.Handler(router.NewRouter()),
		web.RegisterTTL(time.Second*30),
		web.RegisterInterval(time.Second*15),
		web.Metadata(map[string]string{"protocol": "http"}),
	)
	//micro的用法
	_ = webService.Init()
	//将micro中定义的服务与我们自己写的服务绑定在一起
	_ = webService.Run()
}
