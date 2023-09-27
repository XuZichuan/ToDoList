package rpc

import (
	"ToDoList/proto/pb"
	"go-micro.dev/v4"
)

var UserService pb.UserService
var TaskService pb.TaskService

//创建客户端
func InitRPC() {
	userMicroService := micro.NewService(micro.Name("userService.client"))
	//调用实例
	userService := pb.NewUserService("rpcUserService", userMicroService.Client())
	UserService = userService

	taskMicroService := micro.NewService(micro.Name("taskService.client"))
	//调用实例
	taskService := pb.NewTaskService("rpcTaskService", taskMicroService.Client())
	TaskService = taskService
}
