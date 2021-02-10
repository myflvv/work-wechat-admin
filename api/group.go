package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
	"work-wechat-admin/model"
	"work-wechat-admin/utils"
)

func CreateGroup(c *gin.Context) {
	var p model.Group
	err := utils.TranslateZhError(c.ShouldBind(&p))
	if err != nil {
		utils.RespJsonError(http.StatusBadRequest, utils.Error_LackParams, err, c)
		return
	}
	err = p.Create()
	if err != nil {
		utils.RespJsonError(http.StatusInternalServerError, utils.Error_Server, err, c)
		return
	}
	utils.RespJsonOk("success", nil, c)
}

func UpdateGroup(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if id == 0 {
		utils.RespJsonError(http.StatusBadRequest, utils.Error_LackParams, errors.New("缺少ID"), c)
		return
	}
	var p model.Group
	err := utils.TranslateZhError(c.ShouldBind(&p))
	if err != nil {
		utils.RespJsonError(http.StatusBadRequest, utils.Error_LackParams, err, c)
		return
	}
	p.ID = uint(id)
	r, err := p.Update()
	if err != nil {
		utils.RespJsonError(http.StatusInternalServerError, utils.Error_Server, err, c)
		return
	}
	utils.RespJsonOk("success", r, c)
}

func DeleteGroup(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if id == 0 {
		utils.RespJsonError(http.StatusBadRequest, utils.Error_LackParams, errors.New("缺少ID"), c)
		return
	}
	var p model.Group
	p.ID = uint(id)
	err := p.Delete()
	if err != nil {
		utils.RespJsonError(http.StatusInternalServerError, utils.Error_Server, err, c)
		return
	}
	utils.RespJsonOk("success", nil, c)
}

func DetailGroup(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if id == 0 {
		utils.RespJsonError(http.StatusBadRequest, utils.Error_LackParams, errors.New("缺少ID"), c)
		return
	}
	var p model.Group
	p.ID = uint(id)
	r, err := p.Detail()
	if err != nil {
		utils.RespJsonError(http.StatusInternalServerError, utils.Error_Server, err, c)
		return
	}
	utils.RespJsonOk("success", r, c)
}

type selectResp struct {
	Page     int           `json:"page"`
	PageSize int           `json:"page_size"`
	Total    int64           `json:"total"`
	TotalPage int `json:"total_page"`
	Result   []model.Group `json:"result"`
}


func SelectGroup(c *gin.Context) {
	var groups selectResp
	groups.Page, _ = strconv.Atoi(c.Query("page"))
	groups.PageSize, _ = strconv.Atoi(c.Query("page_size"))
	if groups.Page==0 {
		groups.Page=1
	}
	if groups.PageSize>30 || groups.PageSize==0 {
		groups.PageSize=10
	}
	model.DB.Model(&groups.Result).Count(&groups.Total)
	model.DB.Limit(groups.PageSize).Offset((groups.Page - 1) * groups.PageSize).Find(&groups.Result)
	for k,v:=range groups.Result{
		s:=time.Unix(v.CreatedAt,0).Format("2006-01-02 15:04:05")
		groups.Result[k].FormatCreatedAt=s
	}
	utils.RespJsonOk("", groups, c)
}
