package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Resp struct {
	Code int `json:"code"`
	Msg string `json:"msg,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

const (
	Error_LackParams = 40001 //缺少请求参数
	Error_BadParams = 40002 //请求参数错误

	Error_Server = 50000 //内部服务器错误
)

func RespJsonOk(msg string,data interface{},c *gin.Context)  {
	c.JSON(http.StatusOK,Resp{
		Code:0,
		Msg:msg,
		Data:data,
	})
	return
}
func RespJsonError(httpcode int,code int,err error,c *gin.Context)  {
	c.JSON(httpcode,Resp{
		Code:code,
		Msg:err.Error(),
	})
	return
}