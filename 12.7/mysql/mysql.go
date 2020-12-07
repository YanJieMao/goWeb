package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db,err := sql.Open("mysql","root:1997@(localhost:3306)/test?charset=utf8")
	if err != nil {
		panic(err)
	}

	fmt.Println(db.Ping())
	 err = db.Ping()
	if err != nil{
        fmt.Println("数据库链接失败")
	} 
	
	
	
	

	defer db.Close()

	rows,_ := db.Query("select * from blacklist")//查询
	var id int
	var name string

	for rows.Next() {
		rows.Scan(&id,&name)

		fmt.Println(id)
		fmt.Println(name)

	} 
	
}