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
		v.POST("group",api.CreateGroup)
		v.PUT("group/:id",api.UpdateGroup)
		v.DELETE("group/:id",api.DeleteGroup)
		v.GET("group/:id",api.DetailGroup)
		v.GET("group",api.ListGroup)

		v.POST("role",api.CreateRole)
		v.PUT("role/:id",api.UpdateRole)
		v.DELETE("role/:id",api.DeleteRole)
		v.GET("role/:id",api.DetailRole)
		v.GET("role",api.ListRole)
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