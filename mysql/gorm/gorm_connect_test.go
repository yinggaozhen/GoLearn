package gorm

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/yinggaozhen/GoLearn/mysql/stub"
	"testing"
)

// - 官方文档 : {@link https://gorm.io/docs}
// - github地址 : {@link https://github.com/jinzhu/gorm}

func TestConnect(t *testing.T) {
	var dsn = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", stub.User, stub.Password, stub.Database)

	connection, err := gorm.Open("mysql", dsn)
	defer connection.Close()

	if err != nil {
		panic(err)
	}
}

func TestErrorConnect(t *testing.T) {
	var dsn = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", stub.User, "123456", stub.Database)

	connection, err := gorm.Open("mysql", dsn)
	defer connection.Close()

	if err != nil {
		panic(err)
	}
}