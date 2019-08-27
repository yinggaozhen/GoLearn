package sqldriver

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/yinggaozhen/GoLearn/mysql/stub"
	"testing"
)

func TestSelect(t *testing.T) {
	stub.TestSetup()

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", stub.User, stub.Password, stub.Database))
	if err != nil {
		panic(err)
	}

	result, err := db.Query("SELECT * FROM test.salary")
	if err != nil {
		panic(err)
	}

	var id int
	var salary int
	var people string

	for result.Next() == true {
		err := result.Scan(&id, &salary, &people)
		if err != nil {
			panic(err)
		}

		fmt.Println(id, salary, people)
	}
}

func TestInsert(t *testing.T) {
	stub.TestSetup()

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", stub.User, stub.Password, stub.Database))
	if err != nil {
		panic(err)
	}

	result, err := db.Exec("INSERT INTO test.salary (`id`, `salary`, `people`) VALUES (5, 10086, 'jerry')")
	if err != nil {
		panic(err)
	}

	fmt.Println(result.LastInsertId())
}

func TestUpdate(t *testing.T) {
	stub.TestSetup()

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", stub.User, stub.Password, stub.Database))
	if err != nil {
		panic(err)
	}

	result, err := db.Exec("UPDATE test.salary SET salary = 10001 WHERE people = 'jerry'")
	if err != nil {
		panic(err)
	}

	fmt.Println(result.RowsAffected())
}

func TestDelete(t *testing.T) {
	stub.TestSetup()

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", stub.User, stub.Password, stub.Database))
	if err != nil {
		panic(err)
	}

	result, err := db.Exec("DELETE FROM test.salary WHERE id = 5")
	if err != nil {
		panic(err)
	}

	fmt.Println(result.RowsAffected())
}