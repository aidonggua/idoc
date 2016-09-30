package service

import (
	"gorm"
	"time"
)

type Doc struct {
	Id         int       `field:"id"`
	CreateTime time.Time `field:"create_time"`
	ModifyTime time.Time `field:"modify_time"`
	Status     int       `field:"status"`
	ModuleId   int       `field:"module_id"`
	Name       string    `field:"name"`
	Md         string    `field:"md"`
	Sn         int       `field:"sn"`
}

func (d *Doc) GetTableName() string {
	return "tb_doc"
}

//查询所有文档
func QueryAllDoc() []Doc {
	docs := new([]Doc)
	gorm.QueryAll(docs)
	return *docs
}

//保存文档
func SaveDoc(d *Doc) (int64, error) {
	return gorm.Save(d)
}

//删除文档
func DeleteDoc(id int) (int64, error) {
	return gorm.Delete(&Doc{Id: id})
}
