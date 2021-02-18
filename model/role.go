package model

import "time"

type Role struct {
	DefaultField
	Name string `json:"name" form:"name" binding:"required" label:"名称"`
	Tag string	`json:"tag" form:"tag" binding:"required" label:"标识"`
	Remark string `json:"remark" form:"name"`
	Disable *int `json:"disable" gorm:"type:tinyint(1);default:0;comment:是否禁用" form:"disable" label:"是否禁用"`
	GroupId *int `json:"group_id" gorm:"type:int(10);default:0" form:"group_id" label:"组ID"`
}

//详情及查询返回结构体
type RoleResp struct {
	Role
	FormatCreatedAt string `json:"format_create_time"`
	FormatUpdatedAt string `json:"format_updated_time"`
}

func (p *Role) Create() error {
	r := DB.Create(p)
	return r.Error
}

func (p *Role) Update() (*Role, error) {
	r := DB.Model(&p).Updates(Role{Name: p.Name,Tag:p.Tag,Remark:p.Remark,GroupId:p.GroupId,Disable: p.Disable})
	return p, r.Error
}

func (p *Role) Delete() error {
	r := DB.Delete(&p)
	return r.Error
}

func (p *Role) Detail() (*RoleResp, error) {
	r := DB.First(&p)
	var resp RoleResp
	resp.Role=*p
	resp.FormatCreatedAt = time.Unix(p.CreatedAt,0).Format("2006-01-02 15:04:05")
	resp.FormatUpdatedAt= time.Unix(p.UpdatedAt,0).Format("2006-01-02 15:04:05")
	return &resp, r.Error
}