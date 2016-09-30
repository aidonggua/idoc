package service

import (
	"gorm"
	"time"
)

type Module struct {
	Id         int       `field:"id"`
	CreateTime time.Time `field:"create_time"`
	ModifyTime time.Time `field:"modify_time"`
	Status     int       `field:"status"`
	Name       string    `field:"name"`
	AppId      int       `field:"app_id"`
	Sn         int       `field:"sn"`
}

func (m *Module) GetTableName() string {
	return "tb_module"
}

//查询所有模块
func QueryAllModule() []Module {
	modules := new([]Module)
	gorm.QueryAll(modules)
	return *modules
}

//保存模块
func SaveModule(m *Module) (int64, error) {
	return gorm.Save(m)
}

//删除模块
func DeleteModule(id int) (int64, error) {
	return gorm.Delete(&Module{Id: id})
}
