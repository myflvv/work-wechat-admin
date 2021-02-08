package model

type Role struct {
	DefaultField
	Name string `json:"name"`
	Tag string	`json:"tag"`
	Remark string `json:"remark"`
}