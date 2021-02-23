package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"work-wechat-admin/api"
	"work-wechat-admin/model"
	"work-wechat-admin/utils"
)

func main()  {
	_= model.DB.AutoMigrate(
		&model.User{},
		&model.Role{},
		&model.Permission{},
		&model.RolePermissionRelation{},
		&model.Group{},
		)
	s:=utils.Claims{UserId:222,RoleId:333}
	e,_:=s.CreateToken()
	log.Println(e)
	//
	//r,_:=utils.ValidateTokens(e)
	//rr:=r.(map[string]interface{})
	//log.Println(rr["role_id"])
	router := gin.Default()
	v := router.Group("api/v1")
	{
		v.POST("login",api.Login)
	}
	v.Use(jwtMiddleware())
	{
		//其他服务验证jwt,返回claims中的用户信息
		v.POST("verifyjwt",api.VerifyJwt)

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

		v.POST("permission",api.CreatePermission)
		v.PUT("permission/:id",api.UpdatePermission)
		v.DELETE("permission/:id",api.DeletePermission)
		v.GET("permission",api.ListPermission)

		v.POST("relation",api.CreateRelation)
		v.PUT("relation/:id",api.UpdateRelation)
		v.DELETE("relation/:id",api.DeleteRelation)
		v.GET("relation",api.DetailRelation)
		//v.GET("relation",api.ListRelation)
	}
	log.Fatal(router.Run(":8080"))
}

func jwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		Authorization:=c.Request.Header.Get("Authorization")
		if Authorization==""{
			utils.RespJsonError(http.StatusUnauthorized, utils.Error_Authorized, errors.New("token不存在"), c)
			c.Abort()
			return
		}
		token := Authorization[len(BEARER_SCHEMA):]
		res,err:=utils.ValidateToken(token)
		if err!=nil {
			utils.RespJsonError(http.StatusUnauthorized, utils.Error_Authorized, err, c)
			c.Abort()
			return
		}
		c.Set("claims",res)
		c.Next()
	}
}