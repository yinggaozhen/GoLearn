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

// 更新所有记录
func TestUpdateAll(t *testing.T) {
	stub.TestSetup()

	var connection = GetConnection()

	newRecord := model.Salary{
		Salary: 10086,
	}

	db := connection.
		Table("salary").
		Updates(newRecord)

	fmt.Println(db.RowsAffected)
}

// 更新指定记录
func TestUpdateById(t *testing.T) {
	stub.TestSetup()

	var connection = GetConnection()

	newRecord := model.Salary{
		Salary: 10086,
	}

	db := connection.
		Table("salary").
		Where("id = ?", 4).
		Updates(newRecord)

	fmt.Println(db.RowsAffected)
}

// 1. 测试更新注入
func TestUpdateInject(t *testing.T) {
	stub.TestSetup()

	var connection = GetConnection()

	newRecord := model.Salary{
		Salary: 10086,
		People: "lisa';DELETE FROM salary;",
	}

	db := connection.
		Debug().
		Table("salary").
		Where("id = ?", 4).
		Updates(newRecord)

	fmt.Println(db.RowsAffected)
}

// 2. 测试更新注入2
func TestUpdateInject2(t *testing.T) {
	stub.TestSetup()

	var connection = GetConnection()

	newRecord := map[string]interface{}{
		"Salary" : 10086,
		"People" : "lisa';DELETE FROM salary;",
	}

	db := connection.
		Debug().
		Table("salary").
		Where("id = ?", 4).
		Updates(newRecord)

	fmt.Println(db.RowsAffected)
}