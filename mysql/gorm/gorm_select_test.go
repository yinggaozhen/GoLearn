package gorm

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"testing"
)

// - 官方文档 : {@link https://gorm.io/docs}
// - github地址 : {@link https://github.com/jinzhu/gorm}

// 查询所有
func TestSelectAll(t *testing.T) {
	var connection = GetConnection()
	var salaries []Salary

	// SELECT * FROM test.salary WHERE id > 0
	connection.
		Table("salary").
		Find(&salaries, "id > 0")

	fmt.Println(salaries)
}

// 按照ID来查询
func TestSelectById(t *testing.T) {
	var connection = GetConnection()
	var salaries []Salary

	// SELECT * FROM test.salary WHERE id = 1
	connection.
		Table("salary").
		Where("id = 1").
		Find(&salaries)

	fmt.Println(salaries)
}

// 按照IDs来查询(参数绑定)
func TestSelectPrepareById(t *testing.T) {
	var connection = GetConnection()
	var salaries []Salary

	// SELECT * FROM test.salary WHERE id IN (2, 3)
	connection.
		Table("salary").
		Where("id IN (?)", []int{2, 3}).
		Find(&salaries)

	fmt.Println(salaries)
}

// 原生SQL查询
func TestSelectByRawSQL(t *testing.T) {
	var connection = GetConnection()
	var salaries []Salary

	connection.Raw("SELECT * FROM salary WHERE id = 4").Scan(&salaries)

	fmt.Println(salaries)
}

// 获取执行的sql
func TestGetExecSql(t *testing.T) {
	var connection = GetConnection()
	var salaries []Salary

	// SELECT * FROM test.salary WHERE id IN (2, 3)
	sql := connection.
		Table("salary").
		Where("id IN (?)", []int{2, 3}).
		Find(&salaries).
		SubQuery()

	fmt.Println(sql)
}