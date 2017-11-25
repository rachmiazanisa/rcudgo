package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	"fmt"
)

func main() {
	db, err := sql.Open("mysql", "root:@/homebase")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	//Inserting Data
	stmtIns, err := db.Prepare("INSERT INTO user VALUES('','rachmi',23)")
	if err != nil {
		panic(err.Error())
	}
	defer stmtIns.Close()

	_, err = stmtIns.Exec()
	if err != nil {
		panic(err.Error())
	}

	//Deleting Data
	stmtDel, err := db.Prepare("DELETE FROM user WHERE id = ?")
	if err != nil{
		panic(err.Error())
	}
	defer stmtDel.Close()
	_,err = stmtDel.Exec(2)
	if err != nil {
		panic(err.Error())
	}

	//Updating Data
	stmtUpd, err := db.Prepare("UPDATE user SET nama= ? WHERE id = ?")
	if err != nil{
		panic(err.Error())
	}
	defer stmtUpd.Close()
	_,err = stmtUpd.Exec("mamat",3)

	//Reading Data
	stmtOut, err := db.Prepare("SELECT nama FROM user WHERE id = ? ")
	if err != nil {
		panic(err.Error())
	}
	defer stmtOut.Close()

	//Scan the result
	var name string

	err = stmtOut.QueryRow(1).Scan(&name)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("The name  of 1 is: %s", name)

}
