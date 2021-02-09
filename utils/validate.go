package utils

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"net/http"
	"reflect"
	"strings"
)

var trans ut.Translator

func init() {
	uni := ut.New(zh.New())
	trans,_=uni.GetTranslator("zh")
	if v,ok:=binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(field reflect.StructField) string {
			name := strings.SplitN(field.Tag.Get("label"),",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
		//自定义验证
		//_=v.RegisterValidation("logincheck",app.LoginFormatCheck)
		_ = zh_translations.RegisterDefaultTranslations(v,trans)
		return
	}
	return
}

func Verify(request interface{},c *gin.Context) error {
	err:=c.ShouldBind(request)
	if err!=nil {
		var msg interface{}
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			msg = err.Error()
		}else {
			msg = removeTostring(errs.Translate(trans))
		}
		c.JSON(http.StatusBadRequest,gin.H{"code":1,"msg":msg})
		return err
	}
	return nil
}

//翻译为中文错误
func TranslateZhError(err error) error {
	if err != nil {
		var msg string
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			msg = err.Error()
		}else {
			msg = removeTostring(errs.Translate(trans))
		}
		return errors.New(msg)
	}
	return nil
}

//转换map为string,只显示第一个错误信息
func removeTostring(fields map[string]string) string {
	//只显示第一个错误
	i:=0
	var res string
	for _, err := range fields {
		if i==0 {
			res=err
		}
		i++
	}
	return res
}
