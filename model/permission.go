package model

type Permission struct {
	DefaultField
	Title string `json:"title"`
	Method string `json:"method"`
	Path string `json:"path"`
	Pid int `json:"pid" gorm:"type:int(10);default:0"`
	Validate int `json:"validate" gorm:"type:tinyint(1);default:0;comment:是否验证"`
	GroupId int `json:"group_id" grom:"type:int(10)"`
} 
