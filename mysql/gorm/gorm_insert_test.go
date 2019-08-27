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

// 插入单条数据
func TestInsert(t *testing.T) {
	stub.TestSetup()

	var connection = GetConnection()

	// 1. 新增 <熊大> 记录
	xiongda := &model.Salary{
		People: "熊大",
		Salary: 10086,
	}
	connection.
		Table("salary").
		Save(xiongda)
	// 输出自增主键值
	fmt.Println(xiongda.Id)

	// 2. 新增 <熊二> 记录
	xionger := &model.Salary{
		People: "熊二",
		Salary: 10087,
	}
	connection.
		Table("salary").
		Save(xionger)
	// 输出自增主键值
	fmt.Println(xionger.Id)
}

// 测试注入
func TestInsertInject(t *testing.T) {
	stub.TestSetup()

	var connection = GetConnection()

	// 1. 正常新增 <熊大> 记录
	xiongda := &model.Salary{
		People: "熊大",
		Salary: 10086,
	}
	connection.
		Debug().
		Table("salary").
		Save(xiongda)

	// 输出自增主键值
	fmt.Println(xiongda.Id)

	// FIXME 2. 尝试注入 <熊二> 记录
	xionger := &model.Salary{
		People: "熊二');DELETE FROM salary WHERE id = 1;",
		Salary: 10087,
	}
	connection.
		Debug().
		Table("salary").
		Save(xionger)

	// 输出自增主键值
	fmt.Println(xionger.Id)
}