package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm"
)

func main() {
	fmt.Println("hello world")

	gorm.InitDB("mysql", "root:root@tcp(127.0.0.1:3306)/idoc")
	str, err := gorm.Generate("tb_app")
	if err != nil {
		panic(err)
	}
	fmt.Println(str)
}
