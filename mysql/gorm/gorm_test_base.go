package gorm

import (
	"fmt"
	"github.com/jinzhu/gorm"
	gm "github.com/yinggaozhen/GoLearn/mysql"
	"sync"
)

var once sync.Once
var connection *gorm.DB

type Salary struct {
	Id int `gorm:"column:id"`
	Salary int `gorm:"column:salary"`
	People string `gorm:"column:people"`
}

func GetConnection() *gorm.DB {
	once.Do(func() {
		var dsn = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", gm.User, gm.Password, gm.Database)

		connection, _ = gorm.Open("mysql", dsn)
	})

	return connection
}
