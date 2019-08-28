package sqldriver

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/yinggaozhen/GoLearn/mysql/stub"
	"testing"
)

func TestPrepareSelect(t *testing.T) {
	stub.TestSetup()

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", stub.User, stub.Password, stub.Database))
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("SELECT * FROM test.salary WHERE people = ?")
	if err != nil {
		panic(err)
	}

	var id int
	var salary int
	var people string

	result, err := stmt.Query("'0';DELETE FROM salary;")
	for result.Next() == true {
		err := result.Scan(&id, &salary, &people)
		if err != nil {
			panic(err)
		}

		fmt.Println(id, salary, people)
	}
}

func TestPrepareInsert(t *testing.T) {
	stub.TestSetup()

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", stub.User, stub.Password, stub.Database))
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("INSERT INTO test.salary (`id`, `salary`, `people`) VALUES (?, ?, ?)")
	result, err := stmt.Exec(5, 10086, "'');DELETE * FROM salary;")
	if err != nil {
		panic(err)
	}

	fmt.Println(result.LastInsertId())
}


func TestPrepareUpdate(t *testing.T) {
	stub.TestSetup()

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", stub.User, stub.Password, stub.Database))
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("UPDATE test.salary SET salary = ? WHERE people = ?")
	result, err := stmt.Exec(10001, "jerry")
	if err != nil {
		panic(err)
	}

	fmt.Println(result.RowsAffected())
}

func TestPrepareDelete(t *testing.T) {
	stub.TestSetup()

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", stub.User, stub.Password, stub.Database))
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("DELETE FROM test.salary WHERE id = ?")
	result, err := stmt.Exec(5)
	if err != nil {
		panic(err)
	}

	fmt.Println(result.RowsAffected())
}