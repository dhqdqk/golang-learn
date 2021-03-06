package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

/*
CREATE TABLE `userinfo` (
    `uid` INT(10) NOT NULL AUTO_INCREMENT,
    `username` VARCHAR(64) NULL DEFAULT NULL,
    `departname` VARCHAR(64) NULL DEFAULT NULL,
    `created` DATE NULL DEFAULT NULL,
    PRIMARY KEY (`uid`)
)

CREATE TABLE `userdetail` (
    `uid` INT(10) NOT NULL DEFAULT `0`,
    `intro` TEXT NULL,
    `profile` TEXT NULL,
    PRIMARY KEY (`uid`)
)
*/

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// 数据库test，用户表userinfo，关联用户信息表userdetail
func main() {
	// 连接数据库sql.Open("sqlDriveName", "username:password@[tcp([host]:80)]/dbname[?charset=utf8]")
	db, err := sql.Open("mysql", "root:mysql@/test?charset=utf8")
	checkErr(err)

	// 插入数据
	stmt, err := db.Prepare("INSERT userinfo SET username=?,departname=?,created=?")
	checkErr(err)

	res, err := stmt.Exec("golang", "mysql", "2016-05-15")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)

	// 更新数据
	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)

	res, err = stmt.Exec("golang", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)

	// 查询数据
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}

	// 删除数据
	stmt, err = db.Prepare("delete from userinfo where uid=?")
	checkErr(err)

	res, err = stmt.Exec(id)
	checkErr(err)

	affect, err = res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	db.Close()
}
