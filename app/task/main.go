package main

import (
	"ToDoList/app/task/script"
	log "ToDoList/logger"
	"context"
	"fmt"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"

	"ToDoList/app/task/repo/db/dao"
	"ToDoList/app/task/repo/mq"
	"ToDoList/app/task/service"
	"ToDoList/config"
	"ToDoList/proto/pb"
)

func main() {
	config.Init()
	dao.InitDB()
	log.InitLog()
	mq.InitRabbitMQ()
	loadingScript()
	//
	etcdReg := registry.NewRegistry(registry.Addrs(fmt.Sprintf("%s:%s", config.EtcdHost, config.EtcdPort)))
	microTaskService := micro.NewService(
		micro.Name("rpcTaskService"),
		micro.Address(config.TaskServiceAddress),
		micro.Registry(etcdReg),
	)
	microTaskService.Init()
	//绑定
	pb.RegisterTaskServiceHandler(microTaskService.Server(), service.GetTaskSvc())
	//开启
	_ = microTaskService.Run()

}

func loadingScript() {
	ctx := context.Background()
	go script.RunTaskSycn(ctx)
}
