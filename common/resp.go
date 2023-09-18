package common

import (
	"ToDoList/consts"
	"github.com/gin-gonic/gin"
)

// Response 基础序列化器
type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Error  string      `json:"error"`
}

// RespSuccess 带data成功返回
func RespSuccess(ctx *gin.Context, data interface{}, code ...int) *Response {
	status := consts.SUCCESS
	if code != nil {
		status = code[0]
	}

	if data == nil {
		data = "操作成功"
	}

	r := &Response{
		Status: status,
		Data:   data,
		Msg:    GetMsg(status),
	}

	return r
}

func RespError(ctx *gin.Context, err error, data string, code ...int) *Response {
	status := consts.ERROR
	if code != nil {
		status = code[0]
	}

	r := &Response{
		Status: status,
		Data:   data,
		Msg:    GetMsg(status),
		Error:  err.Error(),
	}

	return r
}

var MsgFlags = map[int]string{
	consts.SUCCESS:       "ok",
	consts.ERROR:         "fail",
	consts.InvalidParams: "请求参数错误",
}

// GetMsg 获取状态码对应信息
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[consts.ERROR]
}
