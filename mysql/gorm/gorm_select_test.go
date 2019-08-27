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

// 查询所有
func TestSelectAll(t *testing.T) {
	stub.TestSetup()

	var connection = GetConnection()
	var salaries []model.Salary

	// SELECT * FROM test.salary WHERE id > 0
	connection.
		Table("salary").
		Find(&salaries, "id > 0")

	fmt.Println(salaries)
}

// 按照ID来查询
func TestSelectById(t *testing.T) {
	stub.TestSetup()

	var connection = GetConnection()
	var salaries []model.Salary

	// SELECT * FROM test.salary WHERE id = 1
	connection.
		Table("salary").
		Where("id = 1").
		Find(&salaries)

	fmt.Println(salaries)
}

// 按照IDs来查询(参数绑定)
func TestSelectPrepareById(t *testing.T) {
	stub.TestSetup()

	var connection = GetConnection()
	var salaries []model.Salary

	// SELECT * FROM test.salary WHERE id IN (2, 3)
	connection.
		Table("salary").
		Where("id IN (?)", []int{2, 3}).
		Find(&salaries)

	fmt.Println(salaries)
}

// 原生SQL查询
func TestSelectByRawSQL(t *testing.T) {
	stub.TestSetup()

	var connection = GetConnection()
	var salaries []model.Salary

	connection.Raw("SELECT * FROM salary WHERE id = 4").Scan(&salaries)

	fmt.Println(salaries)
}

// 查看执行的sql
func TestGetExecSql(t *testing.T) {
	stub.TestSetup()

	var connection = GetConnection()
	var salaries []model.Salary

	// SELECT * FROM test.salary WHERE id IN (2, 3)
	connection.
		Table("salary").
		Debug().
		Where("id IN (?)", []int{2, 3}).
		Find(&salaries)
}