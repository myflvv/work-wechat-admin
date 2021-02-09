package model

type Group struct {
	DefaultField
	Name string `json:"name" gorm:"type:varchar(20)" form:"name" binding:"required" label:"名称"`
	Disable *int `json:"disable" gorm:"type:tinyint(1);default:0" form:"disable" binding:"required" label:"是否禁用"`
}

func (p *Group)Create() error {
	r:=DB.Create(p)
	return r.Error
}

func (p *Group)Update() (*Group,error) {
	r:=DB.Model(&p).Updates(Group{Name:p.Name,Disable:p.Disable})
	return p,r.Error
}

func (p *Group)Delete() error {
	r:=DB.Delete(&p)
	return r.Error
}

func (p *Group)Find() (*Group,error) {
	r:=DB.First(&p)
	return p,r.Error
}