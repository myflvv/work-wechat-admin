package model

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	c "work-wechat-admin/config"
)

var DB *gorm.DB

func init()  {
	DB,_ = gorm.Open(mysql.New(mysql.Config{
		DSN:fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",c.Config.DB.User,c.Config.DB.Pass,c.Config.DB.Host,c.Config.DB.Port,c.Config.DB.Name),
		DefaultStringSize:256,
	}),&gorm.Config{
		NamingStrategy:schema.NamingStrategy{SingularTable:true}, //禁用复数表名
		Logger:logger.Default.LogMode(logger.Info),
	})
}