package handler

import (
	"ToDoList/app/gateway/rpc"
	"ToDoList/common"
	"ToDoList/proto/pb"
	"ToDoList/types"
	utils "ToDoList/utils/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// web登陆鉴权、
func UserLoginHandler(ctx *gin.Context) {
	req := &pb.UserRequest{}
	err := ctx.ShouldBind(req)
	if err != nil {
		ctx.JSON(http.StatusOK, common.RespError(ctx, err, "参数绑定失败"))
		return
	}
	resp, err := rpc.UserLogin(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusOK, common.RespError(ctx, err, "UserLogin RPC调用失败"))
		return
	}
	token, err := utils.GenerateToken(uint(resp.UserDetail.Id))
	if err != nil {
		ctx.JSON(http.StatusOK, common.RespError(ctx, err, "GenerateToken失败"))
		return
	}
	respWithToken := &types.TokenData{
		Token: token,
		User:  resp,
	}

	ctx.JSON(http.StatusOK, common.RespSuccess(ctx, respWithToken))
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
