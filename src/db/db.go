package main

import (
	"fmt"
	//"log"
	//"database/sql"
)

func main() {
	//rows, err := db.Query("");
	fun()
}

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
