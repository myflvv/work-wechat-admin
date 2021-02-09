package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"work-wechat-admin/model"
	"work-wechat-admin/utils"
)


func CreateGroup(c *gin.Context)  {
	var p model.Group
	err := utils.TranslateZhError(c.ShouldBind(&p))
	if err != nil {
		utils.RespJsonError(http.StatusBadRequest,utils.Error_LackParams,err,c)
		return
	}
	err =p.Create()
	if err !=nil {
		utils.RespJsonError(http.StatusInternalServerError,utils.Error_Server,err,c)
		return
	}
	utils.RespJsonOk("success",nil,c)
}

func UpdateGroup(c *gin.Context)  {
	id,_:=strconv.Atoi(c.Param("id"))
	if id == 0 {
		utils.RespJsonError(http.StatusBadRequest,utils.Error_LackParams,errors.New("缺少ID"),c)
		return
	}
	var p model.Group
	err := utils.TranslateZhError(c.ShouldBind(&p))
	if err != nil {
		utils.RespJsonError(http.StatusBadRequest,utils.Error_LackParams,err,c)
		return
	}
	p.ID=uint(id)
	r,err:= p.Update()
	if err !=nil {
		utils.RespJsonError(http.StatusInternalServerError,utils.Error_Server,err,c)
		return
	}
	utils.RespJsonOk("success",r,c)
}

func DeleteGroup(c *gin.Context)  {
	id,_:=strconv.Atoi(c.Param("id"))
	if id == 0 {
		utils.RespJsonError(http.StatusBadRequest,utils.Error_LackParams,errors.New("缺少ID"),c)
		return
	}
	var p model.Group
	p.ID=uint(id)
	err:= p.Delete()
	if err !=nil {
		utils.RespJsonError(http.StatusInternalServerError,utils.Error_Server,err,c)
		return
	}
	utils.RespJsonOk("success",nil,c)
}

func FindGroup(c *gin.Context)  {
	id,_:=strconv.Atoi(c.Param("id"))
	if id == 0 {
		utils.RespJsonError(http.StatusBadRequest,utils.Error_LackParams,errors.New("缺少ID"),c)
		return
	}
	var p model.Group
	p.ID=uint(id)
	r,err:= p.Find()
	if err !=nil {
		utils.RespJsonError(http.StatusInternalServerError,utils.Error_Server,err,c)
		return
	}
	utils.RespJsonOk("success",r,c)
}