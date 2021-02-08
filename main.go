package main

import (
	"log"
	"work-wechat-admin/model"
)

func main()  {
	_=model.DB.AutoMigrate(model.User{},&model.Role{})
	token,err:=model.CreateToken(100)
	if err!=nil {
		log.Println(err)
	}
	log.Println(token)
	//time.Sleep(time.Duration(5)*time.Second)
	id,err:= model.ValidateToken(token)
	if err!=nil {
		log.Println(err)
	}
	log.Println(id)
}