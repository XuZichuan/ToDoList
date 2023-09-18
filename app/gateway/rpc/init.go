package rpc

import (
	"ToDoList/idl/pb"
	"go-micro.dev/v4"
)

var UserService pb.UserService

//创建客户端
func InitRPC() {
	userMicroService := micro.NewService(micro.Name("userService.client"))
	//调用实例
	userService := pb.NewUserService("rpcUserService", userMicroService.Client())
	UserService = userService
}
