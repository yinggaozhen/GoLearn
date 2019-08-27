package gendry

import (
	"database/sql"
	"fmt"
	gendry "github.com/didi/gendry/builder"
	_ "github.com/go-sql-driver/mysql"
	"github.com/yinggaozhen/GoLearn/mysql/stub"
	"testing"
)

// Gendry文档 : {@link https://github.com/didi/gendry}
// GendryBuilder文档 : {@link https://github.com/didi/gendry/tree/master/builder}
func TestGendry(t *testing.T) {
	stub.TestSetup()

	// 1. 通过 gendry 构建出sql
	where := map[string]interface{}{
		"id >" : 1,
		"salary": 300,
	}

	cond, vals, err := gendry.BuildSelect("salary", where, []string{"*"})
	fmt.Println(cond, vals, err)

	// 2. 通过原生的 sql-driver 执行sql
	db, _ := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", stub.User, stub.Password, stub.Database))
	stmt, _ := db.Prepare(cond)

	var id int
	var salary int
	var people string

	result, err := stmt.Query(vals...)
	for result.Next() == true {
		result.Scan(&id, &salary, &people)
		fmt.Println(id, salary, people)
	}
}
