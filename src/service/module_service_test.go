package service

import (
	_ "github.com/go-sql-driver/mysql"
	"gorm"
	"testing"
)

func TestSaveModule(t *testing.T) {
	gorm.InitDB("mysql", "root:root@tcp(localhost:3306)/idoc")
	id, err := SaveModule(&Module{Name: "测试模块", AppId: 1, Sn: 1})
	if err != nil {
		t.Log(err)
	}
	t.Log(id)
}

func TestDeleteModule(t *testing.T) {
	gorm.InitDB("mysql", "root:root@tcp(localhost:3306)/idoc")
	success, err := DeleteModule(1)
	if err != nil {
		t.Log(err)
	}
	t.Log(success)
}

func TestQueryAllModule(t *testing.T) {
	gorm.InitDB("mysql", "root:root@tcp(localhost:3306)/idoc")
	modules := QueryAllModule()
	t.Log(modules)
}
