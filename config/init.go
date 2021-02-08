package config

import (
	"github.com/jinzhu/configor"
	"time"
)

var Config = struct {
	DB struct {
		Name         string
		Host         string
		User         string
		Pass         string
		Port         string        `default:"3306"`
		MaxIdleConns int           `default:8`
		MaxOpenConns int           `default:16`
		MaxLifetime  time.Duration `default:120`
	}
}{}

func init() {
	_= configor.Load(&Config, "config/config.yml")
}