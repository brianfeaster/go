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

func openDatabase () *sql.DB {
	db, err := sql.Open("postgres", "host=localhost dbname=freelogue sslmode=disable")
	if err != nil { fmt.Println("ERROR: sql.Open()", err) }
	return db
}

func createTable (db *sql.DB) {
	var s = "create table data (name char (20) primary key, age integer)"
	stmt, err := db.Prepare(s);
	if err != nil { fmt.Println("ERROR:db.Prepare() ", err) }
	res, err := stmt.Exec();
	if err != nil || res == nil { fmt.Println("ERROR: stmt.Exec()", err) }
	stmt.Close();
}

func addRandomPerson (db *sql.DB) {
	var s = "insert into data values (" + "'brian feaster'" + ", " + "42" + ")";
	fmt.Println(s);
	stmt, err := db.Prepare(s);
	if err != nil { fmt.Println("ERROR:db.Prepare() ", err) }

	res, err := stmt.Exec();
	if err != nil || res == nil { fmt.Println("ERROR: stmt.Exec()", err) }
	stmt.Close();
}

func main() {
	//fun()

	fmt.Println(time.Now());

	//fmt.Println()
	//fmt.Println(db)
	//fmt.Println(err)

	db := openDatabase()

	if db != nil {
		createTable(db)
		addRandomPerson(db)
	}


	db.Close()
}

