package gorm

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/yinggaozhen/GoLearn/mysql/gorm/model"
	"github.com/yinggaozhen/GoLearn/mysql/stub"
	"testing"
)

// - 官方文档 : {@link https://gorm.io/docs}
// - github地址 : {@link https://github.com/jinzhu/gorm}

// 删除所有记录
func TestDeleteAll(t *testing.T) {
	stub.TestSetup()

	var connection = GetConnection()

	db := connection.
		Debug().
		Table("salary").
		Delete(nil)

	fmt.Println(db.RowsAffected)
}

// 根据where删除
func TestDeleteByWhere(t *testing.T) {
	stub.TestSetup()

	var connection = GetConnection()

	db := connection.
		Debug().
		Table("salary").
		Where("id > ?", 2).
		Delete(nil)

	fmt.Println(db.RowsAffected)
}

// 根据model删除
func TestDeleteByModel(t *testing.T) {
	stub.TestSetup()

	var connection = GetConnection()

	deleteRecord := model.Salary{
		Id : 1,
	}

	db := connection.
		Debug().
		Table("salary").
		Delete(deleteRecord)

	fmt.Println(db.RowsAffected)
}