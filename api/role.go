package api

type roleParams struct {
	Name string `form:"name" validate:"required"`
	Tag string `form:"tag" validate:"required"`
	Disable int `form:"disable" validate:"required"`
}