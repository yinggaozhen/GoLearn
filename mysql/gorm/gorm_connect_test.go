package gorm

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	gm "github.com/yinggaozhen/GoLearn/mysql"
	"testing"
)

// - 官方文档 : {@link https://gorm.io/docs}
// - github地址 : {@link https://github.com/jinzhu/gorm}

func TestConnect(t *testing.T) {
	var dsn = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", gm.User, gm.Password, gm.Database)

	connection, err := gorm.Open("mysql", dsn)
	defer connection.Close()

	if err != nil {
		panic(err)
	}
}

func TestErrorConnect(t *testing.T) {
	var dsn = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", gm.User, "123456", gm.Database)

	connection, err := gorm.Open("mysql", dsn)
	defer connection.Close()

	if err != nil {
		panic(err)
	}
}