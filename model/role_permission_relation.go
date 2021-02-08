package model

//角色权限关系
type RolePermissionRelation struct {
	DefaultField
	RoleId int `json:"role_id" gorm:"type:int(10);"`
	PermissionId int `json:"permission_id" gorm:"type:int(10);"`
}  