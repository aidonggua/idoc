package service

import (
	_ "github.com/go-sql-driver/mysql"
	"gorm"
	"testing"
)

func TestSaveApp(t *testing.T) {
	gorm.InitDB("mysql", "root:root@tcp(localhost:3306)/idoc")
	id, err := SaveApp(&App{Name: "测试", Token: "idoc"})
	if err != nil {
		t.Log(err)
	}
	t.Log(id)
}

func TestDeleteApp(t *testing.T) {
	gorm.InitDB("mysql", "root:root@tcp(localhost:3306)/idoc")
	success, err := DeleteApp(&App{Id: 1})
	if err != nil {
		t.Log(err)
	}
	t.Log(success)
}

func TestQueryAllApp(t *testing.T) {
	gorm.InitDB("mysql", "root:root@tcp(localhost:3306)/idoc")
	apps := QueryAllApp()
	t.Log(apps)
}
