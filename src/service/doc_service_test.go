package service

import (
	_ "github.com/go-sql-driver/mysql"
	"gorm"
	"testing"
)

func TestSaveDoc(t *testing.T) {
	gorm.InitDB("mysql", "root:root@tcp(localhost:3306)/idoc")
	id, err := SaveDoc(&Doc{ModuleId: 1, Name: "idoc", Md: "aaa", Sn: 1})
	if err != nil {
		t.Log(err)
	}
	t.Log(id)
}

func TestDeleteDoc(t *testing.T) {
	gorm.InitDB("mysql", "root:root@tcp(localhost:3306)/idoc")
	success, err := DeleteDoc(1)
	if err != nil {
		t.Log(err)
	}
	t.Log(success)
}

func TestQueryAllDoc(t *testing.T) {
	gorm.InitDB("mysql", "root:root@tcp(localhost:3306)/idoc")
	docs := QueryAllDoc()
	t.Log(docs)
}
