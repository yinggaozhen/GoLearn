package gorm

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/yinggaozhen/GoLearn/mysql/stub"
	"sync"
)

var once sync.Once
var connection *gorm.DB

func GetConnection() *gorm.DB {
	once.Do(func() {
		var dsn = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local&multiStatements=true", stub.User, stub.Password, stub.Database)

		connection, _ = gorm.Open("mysql", dsn)
	})

	return connection
}
