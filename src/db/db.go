package main

import (
	"fmt"
	//"log"
	"database/sql"
	_ "github.com/bmizerany/pq"
	"time"
)

func fun() {
	if a := 9; a == 9 {
		fmt.Println("a =", a)
	}

	for i, c := range "abcd" {
		fmt.Println("idx:", i, " char:", c)
	}

	for i, j := 0, 9; i <= j; i, j = i+2, j+1 {
		fmt.Print("[i:", i, " j:", j, "]")
	}
}

func createTable () *DB {
	db, err := sql.Open("postgres", "host=localhost dbname=freelogue sslmode=disable")
	if err != nil { fmt.Println("ERROR: sql.Open()", err) }
	return db
}

func main() {
	//fun()

	//fmt.Println()
	//fmt.Println(db)
	//fmt.Println(err)

	db := createTable()

	var s = "create table data (name char (20) primary key, age integer)"
	stmt, err := db.Prepare(s);
	if err != nil { fmt.Println("ERROR:db.Prepare() ", err) }

	fmt.Println(time.Now());
	res, err := stmt.Exec();
	if err != nil { fmt.Println("ERROR: stmt.Exec()", err) }

	fmt.Println(res);

	stmt.Close();
	db.Close()
}

