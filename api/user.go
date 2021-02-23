package api

import (
	"github.com/gin-gonic/gin"
	"work-wechat-admin/utils"
)

func Login(c *gin.Context)  {

}

//其他应用验证jwt,返回claims中的用户信息
func VerifyJwt(c *gin.Context)  {
	claims:=c.MustGet("claims").(*utils.Claims)
	utils.RespJsonOk("success", claims, c)
}