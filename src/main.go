package main

import (
	"controller"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm"
	"grouter"
	"grow"
)

func main() {
	gorm.InitDB("mysql", "root:root@tcp(localhost:3306)/idoc")
	grouter.Route("/app/list", controller.ListAllApp, "get")
	grow.Start(8888)
}

func generate() {
	gorm.InitDB("mysql", "root:root@tcp(127.0.0.1:3306)/idoc")
	str, err := gorm.Generate("tb_doc")
	if err != nil {
		panic(err)
	}
	fmt.Println(str)
}
