package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestPrepareSelect(t *testing.T) {
	db, err := sql.Open("mysql", "root:@/test")
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("SELECT * FROM test.salary WHERE id > ? AND salary >= ?")
	if err != nil {
		panic(err)
	}

	var id int
	var salary int
	var people string

	result, err := stmt.Query(0, 200)
	for result.Next() == true {
		err := result.Scan(&id, &salary, &people)
		if err != nil {
			panic(err)
		}

		fmt.Println(id, salary, people)
	}
}

func TestPrepareInsert(t *testing.T) {
	db, err := sql.Open("mysql", "root:@/test")
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("INSERT INTO test.salary (`id`, `salary`, `people`) VALUES (?, ?, ?)")
	result, err := stmt.Exec(5, 10086, "jerry")
	if err != nil {
		panic(err)
	}

	fmt.Println(result.LastInsertId())
}


func TestPrepareUpdate(t *testing.T) {
	db, err := sql.Open("mysql", "root:@/test")
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
	db, err := sql.Open("mysql", "root:@/test")
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