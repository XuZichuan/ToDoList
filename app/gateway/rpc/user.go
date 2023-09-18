package rpc

import (
	"ToDoList/consts"
	"ToDoList/idl/pb"
	"github.com/gin-gonic/gin"
)

//
func UserLogin(ctx *gin.Context, request *pb.UserRequest) (*pb.UserDetailResponse, error) {
	resp, err := UserService.UserLogin(ctx, request)
	if err != nil {
		resp.Code = consts.ERROR
		return nil, err
	}
	return resp, nil
}

func UserRegister(ctx *gin.Context, request *pb.UserRequest) (*pb.UserDetailResponse, error) {
	resp, err := UserService.UserRegister(ctx, request)
	if err != nil {
		resp.Code = consts.ERROR
		return nil, err
	}
	return resp, nil
}
