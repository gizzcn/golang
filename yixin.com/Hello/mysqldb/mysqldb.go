package mysqldb

import (
	"database/sql"
	//"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func SqlGet(sqlStr string) string {

	db, err := sql.Open("mysql", "root:123456@/python?charset=utf8")
	CheckErr(err)
	rows, err := db.Query(sqlStr)
	var id int
	var uname string
	if rows.Next() {

		err = rows.Scan(&id, &uname)
		CheckErr(err)

	}
	db.Close()
	return uname
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
