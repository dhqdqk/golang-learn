package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}

}

func main() {
	db, err := sql.Open("sqlite3", "./test")
	checkErr(err)

	// insert data
	stmt, err := db.Prepare("insert into userinfo(username, departname, created) values(?,?,?)")
	checkErr(err)

	res, err := stmt.Exec("golang", "sqlite3", "2016-5-23")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println(id)

	// update data
	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)

	res, err = stmt.Exec("golangupdate", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	// query data
	rows, err := db.Query("select * from userinfo")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created time.Time
		err = rows.Scan(&uid, &username, &department, &created)

	}
}
