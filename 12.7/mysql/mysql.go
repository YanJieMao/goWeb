package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type user struct {
	id int
	name string
}



func queryMultiRow() {

	db,err := sql.Open("mysql","root:1997@(localhost:3306)/test?charset=utf8")
	if err != nil {
		panic(err)
	}

	rows,_ := db.Query("select id,name from blacklist where id < 10")//查询
	
	var u user
	//var name string

	for rows.Next() {
		rows.Scan(&u.id,&u.name)
		fmt.Println(u.id)
		fmt.Println(u.name)
	} 
	
}

func queryAll() {

	db,err := sql.Open("mysql","root:1997@(localhost:3306)/test?charset=utf8")
	if err != nil {
		panic(err)
	}

	rows,_ := db.Query("select * from blacklist")//查询
	var id int
	var name string

	for rows.Next() {
		rows.Scan(&id,&name)

		fmt.Println(id)
		fmt.Println(name)

	} 

}

func insertRow() {

	 db,err := sql.Open("mysql","root:1997@(localhost:3306)/test?charset=utf8")
	if err != nil {
		panic(err)
	} 

	slqstr := "insert into blacklist(name) values (?)"
	ret, err :=db.Exec(slqstr,"aaaa")
	if err != nil {
        fmt.Printf("insert failed, err:%v\n", err)
        return
	}

	theID, err := ret.LastInsertId() // 新插入数据的id
    if err != nil {
        fmt.Printf("get lastinsert ID failed, err:%v\n", err)
        return
    }
    fmt.Printf("insert success, the id is %d.\n", theID)
		
	
}

func updateRow() {

	db,err := sql.Open("mysql","root:1997@(localhost:3306)/test?charset=utf8")
	if err != nil {
		panic(err)
	} 

	sqlstr := "update blacklist set name= ? where id = ?"
	ret, err :=db.Exec(sqlstr,"test",5001)
	if err != nil {
		fmt.Println("update field, err:%v\n", err)
		return
	}

	n, err := ret.RowsAffected() // 操作影响的行数
    if err != nil {
        fmt.Printf("get RowsAffected failed, err:%v\n", err)
        return
    }
    fmt.Printf("update success, affected rows:%d\n", n)



}

func deleteRow() {
	db,err := sql.Open("mysql","root:1997@(localhost:3306)/test?charset=utf8")
	if err != nil {
		panic(err)
	}

	sqlstr := "delete from blacklist where id = ?"
	ret,err := db.Exec(sqlstr,5003)
	if err != nil{
		panic(err)
	}

	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Println("get rows Affected field err:%v|n",err)
		return
	}

	fmt.Printf("delete success, affected rows:%d\n", n)


}



func main() {

	db,err := sql.Open("mysql","root:1997@(localhost:3306)/test?charset=utf8")
	if err != nil {
		panic(err)
	}

	
	 err = db.Ping()
	if err != nil{
        fmt.Println("数据库链接失败")
	} 
	
	deleteRow()
	defer db.Close()
	

	
}