package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestSelect(t *testing.T) {
	db, err := sql.Open("mysql", "root:@/test")
	if err != nil {
		panic(err)
	}

	result, err := db.Query("select * from test.salary")
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
	db, err := sql.Open("mysql", "root:@/test")
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
	db, err := sql.Open("mysql", "root:@/test")
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
	db, err := sql.Open("mysql", "root:@/test")
	if err != nil {
		panic(err)
	}

	result, err := db.Exec("DELETE FROM test.salary WHERE id = 5")
	if err != nil {
		panic(err)
	}

	fmt.Println(result.RowsAffected())
}