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

func CreateRelation(c *gin.Context)  {
	var p model.RolePermissionRelation
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
func UpdateRelation(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if id == 0 {
		utils.RespJsonError(http.StatusBadRequest, utils.Error_LackParams, errors.New("缺少ID"), c)
		return
	}
	var p model.RolePermissionRelation
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

func DeleteRelation(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if id == 0 {
		utils.RespJsonError(http.StatusBadRequest, utils.Error_LackParams, errors.New("缺少ID"), c)
		return
	}
	var p model.RolePermissionRelation
	p.ID = uint(id)
	err := p.Delete()
	if err != nil {
		utils.RespJsonError(http.StatusInternalServerError, utils.Error_Server, err, c)
		return
	}
	utils.RespJsonOk("success", nil, c)
}

func DetailRelation(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if id == 0 {
		utils.RespJsonError(http.StatusBadRequest, utils.Error_LackParams, errors.New("缺少ID"), c)
		return
	}
	var p model.RolePermissionRelation
	p.ID = uint(id)
	r, err := p.Detail()
	if err != nil {
		utils.RespJsonError(http.StatusInternalServerError, utils.Error_Server, err, c)
		return
	}
	utils.RespJsonOk("success", r, c)
}

type selectRolePerRelationResp struct {
	Page     int           `json:"page"`
	PageSize int           `json:"page_size"`
	Total    int64           `json:"total"`
	TotalPage int `json:"total_page"`
	Result   []relationResp `json:"result"`
}

type relationResp struct {
	ID uint `json:"id"`
	CreatedAt int64 `json:"created_at,omitempty"`
	UpdatedAt int64 `json:"updated_at,omitempty"`
	RoleName string `json:"role_name"`
	PermissionTitle string `json:"permission_title"`
	FormatCreatedAt string `json:"format_create_time"`
	FormatUpdatedAt string `json:"format_updated_time"`
} 

//分页
func ListRelation(c *gin.Context) {
	var result selectRolePerRelationResp
	result.Page, _ = strconv.Atoi(c.Query("page"))
	result.PageSize, _ = strconv.Atoi(c.Query("page_size"))
	if result.Page==0 {
		result.Page=1
	}
	if result.PageSize>30 || result.PageSize==0 {
		result.PageSize=10
	}
	model.DB.Table("role_permission_relation").Count(&result.Total)

	var tr []relationResp
	model.DB.Table("role_permission_relation rp").Select("rp.id,r.name as role_name,p.title as permission_title,rp.created_at,rp.updated_at").
		Joins("left join role r on rp.role_id=r.id").
		Joins("left join permission p on rp.permission_id=p.id").
		Limit(result.PageSize).Offset((result.Page - 1) * result.PageSize).Scan(&tr)
	//log.Printf("%+v",tr)
	//return
	result.TotalPage=int(math.Ceil(float64(float64(result.Total)/float64(result.PageSize))))
	for k,v:=range tr{
		tr[k].FormatCreatedAt=time.Unix(v.CreatedAt,0).Format("2006-01-02 15:04:05")
		tr[k].FormatUpdatedAt=time.Unix(v.UpdatedAt,0).Format("2006-01-02 15:04:05")
	}
	result.Result=tr
	utils.RespJsonOk("", result, c)
}