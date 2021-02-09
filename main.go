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

	router := gin.Default()
	v := router.Group("api/v1")
	{
		v.GET("group/:id",api.DetailGroup)
		v.GET("group",api.SelectGroup)
		v.POST("group",api.CreateGroup)
		v.PUT("group/:id",api.UpdateGroup)
		v.DELETE("group/:id",api.DeleteGroup)
	}
	log.Fatal(router.Run(":8080"))
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