package model

type Role struct {
	DefaultField
	Name string `json:"name"`
	Tag string	`json:"tag"`
	Remark string `json:"remark"`
	Disable *int `json:"disable" gorm:"type:tinyint(1);default:0;comment:是否禁用"`
	GroupId *int `json:"group_id" grom:"type:int(10);default:0"`
}

func (r *Role)Insert()  {

}