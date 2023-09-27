package main

import (
	"fmt"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"

	"ToDoList/app/user/repo/db/dao"
	"ToDoList/app/user/service"
	"ToDoList/config"
	"ToDoList/proto/pb"
)

//user微服务
func main() {
	//初始化配置文件等的读取
	config.Init()
	//初始化数据库
	dao.InitDB()
	//注册一个etcd
	etcdReg := registry.NewRegistry(registry.Addrs(fmt.Sprintf("%s:%s", config.EtcdHost, config.EtcdPort)))
	//把当前的服务注册成一个微服务
	microUserService := micro.NewService(
		micro.Name("rpcUserService"),
		micro.Address(config.UserServiceAddress), //服务的地址
		micro.Registry(etcdReg),                  //etcd服务注册
	)
	//micro的用法
	microUserService.Init()
	//将micro中定义的服务与我们自己写的服务绑定在一起
	_ = pb.RegisterUserServiceHandler(microUserService.Server(), service.GetUserSvc())

	_ = microUserService.Run()
}
