package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
	"strconv"
	"time"
	"work-wechat-admin/model"
	"work-wechat-admin/utils"
)

func CreatePermission(c *gin.Context)  {
	var p model.Permission
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

func UpdatePermission(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if id == 0 {
		utils.RespJsonError(http.StatusBadRequest, utils.Error_LackParams, errors.New("缺少ID"), c)
		return
	}
	var p model.Permission
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

func DeletePermission(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if id == 0 {
		utils.RespJsonError(http.StatusBadRequest, utils.Error_LackParams, errors.New("缺少ID"), c)
		return
	}
	var p model.Permission
	p.ID = uint(id)
	err := p.Delete()
	if err != nil {
		utils.RespJsonError(http.StatusInternalServerError, utils.Error_Server, err, c)
		return
	}
	utils.RespJsonOk("success", nil, c)
}

type selectPermissionResp struct {
	Page     int           `json:"page"`
	PageSize int           `json:"page_size"`
	Total    int64           `json:"total"`
	TotalPage int `json:"total_page"`
	Result   []model.PermissionResp `json:"result"`
}

//分页
func ListPermission(c *gin.Context) {
	var result selectPermissionResp
	var resultSlice []model.Permission
	result.Page, _ = strconv.Atoi(c.Query("page"))
	result.PageSize, _ = strconv.Atoi(c.Query("page_size"))
	if result.Page==0 {
		result.Page=1
	}
	if result.PageSize>30 || result.PageSize==0 {
		result.PageSize=10
	}
	model.DB.Model(&resultSlice).Count(&result.Total)
	model.DB.Limit(result.PageSize).Offset((result.Page - 1) * result.PageSize).Find(&resultSlice)
	result.TotalPage=int(math.Ceil(float64(float64(result.Total)/float64(result.PageSize))))
	ef:=make([]model.PermissionResp,len(resultSlice))
	for k,v:=range resultSlice{
		createAt:=time.Unix(v.CreatedAt,0).Format("2006-01-02 15:04:05")
		updateAt:=time.Unix(v.UpdatedAt,0).Format("2006-01-02 15:04:05")
		ef[k].FormatCreatedAt=createAt
		ef[k].FormatUpdatedAt=updateAt
		ef[k].Title=v.Title
		ef[k].Method=v.Method
		ef[k].Path=v.Path
		ef[k].Pid=v.Pid
		ef[k].Validate=v.Validate
		ef[k].GroupId=v.GroupId
		ef[k].ID=v.ID
	}
	result.Result=ef
	utils.RespJsonOk("", result, c)
}