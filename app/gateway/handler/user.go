package handler

import (
	"ToDoList/app/gateway/rpc"
	"ToDoList/common"
	"ToDoList/idl/pb"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserLoginHandler(ctx *gin.Context) {
	var req *pb.UserRequest
	err := ctx.ShouldBind(req)
	if err != nil {
		ctx.JSON(http.StatusOK, common.RespError(ctx, err, "参数绑定失败"))
		return
	}
	resp, err := rpc.UserLogin(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusOK, common.RespError(ctx, err, "UserLogin RPC 调用失败"))
		return
	}
	ctx.JSON(http.StatusOK, common.RespSuccess(ctx, resp))
}

func UserRegisterHandler(ctx *gin.Context) {
	req := &pb.UserRequest{}
	err := ctx.ShouldBind(req)
	if err != nil {
		ctx.JSON(http.StatusOK, common.RespError(ctx, err, "参数绑定失败"))
		return
	}
	resp, err := rpc.UserRegister(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusOK, common.RespError(ctx, err, "UserRegister RPC 调用失败"))
		return
	}
	ctx.JSON(http.StatusOK, common.RespSuccess(ctx, resp))
}
