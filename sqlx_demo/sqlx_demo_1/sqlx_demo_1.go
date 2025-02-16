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

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("inti DB failed,err:%v\n", err)
	}
	defer db.Close()

	fmt.Printf("init DB succeed ……")

}
