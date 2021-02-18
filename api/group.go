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
	Result   []model.GroupResp `json:"result"`
}

//分页
func ListGroup(c *gin.Context) {
	var groups selectResp
	var resultGroup []model.Group
	groups.Page, _ = strconv.Atoi(c.Query("page"))
	groups.PageSize, _ = strconv.Atoi(c.Query("page_size"))
	if groups.Page==0 {
		groups.Page=1
	}
	if groups.PageSize>30 || groups.PageSize==0 {
		groups.PageSize=10
	}
	model.DB.Model(&resultGroup).Count(&groups.Total)
	model.DB.Limit(groups.PageSize).Offset((groups.Page - 1) * groups.PageSize).Find(&resultGroup)
	groups.TotalPage=int(math.Ceil(float64(float64(groups.Total)/float64(groups.PageSize))))
	ef:=make([]model.GroupResp,len(resultGroup))
	for k,v:=range resultGroup{
		createAt:=time.Unix(v.CreatedAt,0).Format("2006-01-02 15:04:05")
		updateAt:=time.Unix(v.UpdatedAt,0).Format("2006-01-02 15:04:05")
		ef[k].FormatCreatedAt=createAt
		ef[k].FormatUpdatedAt=updateAt
		ef[k].Name=v.Name
		ef[k].Disable=v.Disable
		ef[k].ID=v.ID
	}
	groups.Result=ef
	utils.RespJsonOk("", groups, c)
}
