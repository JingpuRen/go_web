package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// 全局变量
var db *sqlx.DB

func initDB() (err error) {
	// tip 注意我们这里要初始化一个全局的db变量，而不是使用:=！！
	db, err = sqlx.Connect("mysql", "root:666666@(localhost:3306)/sqlx_demo?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Printf("connect DB failed,err:%v\n", err)
		return err
	}

	db.SetMaxOpenConns(200)
	db.SetMaxIdleConns(10)

	return nil
}

// 定义User类
type User struct {
	Id   int64  `db:"id"`
	Name string `db:"name"`
	Age  int64  `db:"age"`
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("inti DB failed,err:%v\n", err)
	}
	defer db.Close()

	// todo 单行查询，使用Get方法即可
	var user User
	// tip sql语句中一定要记得使用?来作为占位符！！！
	sqlStr := "select id,name,age from users where id = ?"

	err2 := db.Get(&user, sqlStr, 1)

	if err2 != nil {
		fmt.Printf("the error is : %v\n", err2)
		return
	}

	fmt.Printf("id = %d,name = %s,age = %d\n", user.Id, user.Name, user.Age)

	// todo 多行查询，使用Select方法即可
	var users []User
	sqlStr2 := "select id,name,age from users where id >= ?"
	err3 := db.Select(&users, sqlStr2, 1)
	if err3 != nil {
		fmt.Printf("the error is %v\n", err3)
		return
	}
	for _, value := range users {
		fmt.Printf("id = %d,name = %s,age = %d\n", value.Id, value.Name, value.Age)
	}

	fmt.Printf("the code is exit successfully ......")
}
