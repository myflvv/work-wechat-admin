package model

type Permission struct {
	DefaultField
	Title string `json:"title" form:"title" binding:"required" label:"标题"`
	Method string `json:"method" form:"method" binding:"required" label:"模式"`
	Path string `json:"path" form:"path" binding:"required" label:"路径"`
	Pid *int `json:"pid" gorm:"type:int(10);default:0" form:"pid"`
	Validate *int `json:"validate" gorm:"type:tinyint(1);default:0;comment:是否验证" form:"validate"`
	GroupId *int `json:"group_id" grom:"type:int(10);default:0" form:"group_id" binding:"required,gt=0" label:"组"`
}

//详情及查询返回结构体
type PermissionResp struct {
	Permission
	FormatCreatedAt string `json:"format_create_time"`
	FormatUpdatedAt string `json:"format_updated_time"`
}

func (p *Permission) Create() error {
	r := DB.Create(p)
	return r.Error
}

func (p *Permission) Update() (*Permission, error) {
	r := DB.Model(&p).Updates(Permission{Title: p.Title,Method:p.Method,Path:p.Path,GroupId:p.GroupId,Validate: p.Validate,Pid:p.Pid})
	return p, r.Error
}

func (p *Permission) Delete() error {
	r := DB.Delete(&p)
	return r.Error
}
