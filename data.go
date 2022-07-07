package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func init(){
	dbconf := "myapp:admin@tcp(10.0.2.10:3306)/myapp?charset=utf8mb4"
	// dbconf := "myapp:admin@/myapp"
	var err error
	Db, err = sql.Open("mysql", dbconf)
	if err != nil {
		panic(err)
	}
}

func post(emailaddress string,password string){
	statement := "insert into USER (EMAIL_ADDRESS, PASSWORD) values (?,?)"
	stmt ,err := Db.Prepare(statement)
	if err != nil{
		panic(err)
	}
	defer stmt.Close()
	stmt.Exec(emailaddress,password)
	return
}