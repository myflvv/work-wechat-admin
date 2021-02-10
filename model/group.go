package model

import (
	"time"
)

type Group struct {
	DefaultField
	Name    string `json:"name" gorm:"type:varchar(20)" form:"name" binding:"required" label:"名称"`
	Disable *int   `json:"disable" gorm:"type:tinyint(1);default:0" form:"disable" binding:"required" label:"是否禁用"`
}

type GroupResp struct {
	Group
	FormatCreatedAt string `json:"format_create_time"`
	FormatUpdatedAt string `json:"format_updated_time"`
} 

func (p *Group) Create() error {
	r := DB.Create(p)
	return r.Error
}

func (p *Group) Update() (*Group, error) {
	r := DB.Model(&p).Updates(Group{Name: p.Name, Disable: p.Disable})
	return p, r.Error
}

func (p *Group) Delete() error {
	r := DB.Delete(&p)
	return r.Error
}

func (p *Group) Detail() (*GroupResp, error) {
	r := DB.First(&p)
	var resp GroupResp
	resp.Group=*p
	resp.FormatCreatedAt = time.Unix(p.CreatedAt,0).Format("2006-01-02 15:04:05")
	resp.FormatUpdatedAt= time.Unix(p.UpdatedAt,0).Format("2006-01-02 15:04:05")
	return &resp, r.Error
}