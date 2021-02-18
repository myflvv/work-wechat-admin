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

func CreateRole(c *gin.Context)  {
	var p model.Role
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

func UpdateRole(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if id == 0 {
		utils.RespJsonError(http.StatusBadRequest, utils.Error_LackParams, errors.New("缺少ID"), c)
		return
	}
	var p model.Role
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

func DeleteRole(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if id == 0 {
		utils.RespJsonError(http.StatusBadRequest, utils.Error_LackParams, errors.New("缺少ID"), c)
		return
	}
	var p model.Role
	p.ID = uint(id)
	err := p.Delete()
	if err != nil {
		utils.RespJsonError(http.StatusInternalServerError, utils.Error_Server, err, c)
		return
	}
	utils.RespJsonOk("success", nil, c)
}

func DetailRole(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if id == 0 {
		utils.RespJsonError(http.StatusBadRequest, utils.Error_LackParams, errors.New("缺少ID"), c)
		return
	}
	var p model.Role
	p.ID = uint(id)
	r, err := p.Detail()
	if err != nil {
		utils.RespJsonError(http.StatusInternalServerError, utils.Error_Server, err, c)
		return
	}
	utils.RespJsonOk("success", r, c)
}

type selectRoleResp struct {
	Page     int           `json:"page"`
	PageSize int           `json:"page_size"`
	Total    int64           `json:"total"`
	TotalPage int `json:"total_page"`
	Result   []model.RoleResp `json:"result"`
}

//分页
func ListRole(c *gin.Context) {
	var result selectRoleResp
	var resultSlice []model.Role
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
	ef:=make([]model.RoleResp,len(resultSlice))
	for k,v:=range resultSlice{
		createAt:=time.Unix(v.CreatedAt,0).Format("2006-01-02 15:04:05")
		updateAt:=time.Unix(v.UpdatedAt,0).Format("2006-01-02 15:04:05")
		ef[k].FormatCreatedAt=createAt
		ef[k].FormatUpdatedAt=updateAt
		ef[k].Name=v.Name
		ef[k].Disable=v.Disable
		ef[k].ID=v.ID
	}
	result.Result=ef
	utils.RespJsonOk("", result, c)
}