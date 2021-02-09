package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"work-wechat-admin/api"
	"work-wechat-admin/model"
)

func main()  {
	_= model.DB.AutoMigrate(
		&model.User{},
		&model.Role{},
		&model.Permission{},
		&model.RolePermissionRelation{},
		&model.Group{},
		)

	r := gin.Default()
	verRouter := r.Group("api/v1")
	{
		verRouter.GET("group/:id",api.FindGroup)
		verRouter.POST("group",api.CreateGroup)
		verRouter.PUT("group/:id",api.UpdateGroup)
		verRouter.DELETE("group/:id",api.DeleteGroup)
	}
	log.Fatal(r.Run(":8080"))
	//token,err:=model.CreateToken(100)
	//if err!=nil {
	//	log.Println(err)
	//}
	//log.Println(token)
	//id,err:= model.ValidateToken(token)
	//if err!=nil {
	//	log.Println(err)
	//}
	//log.Println(id)
}