package model

import "gorm.io/gorm"

type DefaultField struct {
	ID uint `json:"id" gorm:"primarykey"`
	CreatedAt int64 `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt int64 `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type User struct {
	DefaultField
	WechatUserId string `json:"wechat_user_id" gorm:"type:varchar(50)"` //微信userid
	Username string `json:"username" gorm:"type:varchar(50)"`
	Password string `json:"password" gorm:"type:varchar(50)"`
	RoleId int `json:"role_id" gorm:"type:int(10)"`
}