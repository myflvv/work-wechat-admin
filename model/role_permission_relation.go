package model

//角色权限关系
type RolePermissionRelation struct {
	DefaultField
	RoleId int `json:"role_id" gorm:"type:int(10);" form:"role_id" binding:"required,gt=0" label:"角色ID"`
	PermissionId string `json:"permission_id" gorm:"type:text;" form:"permission_id" binding:"required" label:"权限ID"`
}

func (p *RolePermissionRelation) Create() error {
	r := DB.Create(p)
	return r.Error
}
func (p *RolePermissionRelation) Update() (*RolePermissionRelation, error) {
	r := DB.Model(&p).Updates(RolePermissionRelation{PermissionId:p.PermissionId})
	return p, r.Error
}

func (p *RolePermissionRelation) Detail() (*RolePermissionRelation, error) {
	r := DB.First(&p)
	return p, r.Error
}

func (p *RolePermissionRelation) Delete() error {
	r := DB.Delete(&p)
	return r.Error
}