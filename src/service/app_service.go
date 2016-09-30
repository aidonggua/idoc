package service

import (
	"gorm"
	"time"
)

type App struct {
	Id         int       `field:"id"`
	CreateTime time.Time `field:"create_time"`
	ModifyTime time.Time `field:"modify_time"`
	Status     int       `field:"status"`
	Name       string    `field:"name"`
	Token      string    `field:"token"`
}

func (a *App) GetTableName() string {
	return "tb_app"
}

//保存App
func SaveApp(app *App) (int64, error) {
	app.ModifyTime = time.Now()
	return gorm.Save(app)
}

//删除App
func DeleteApp(id int) (int64, error) {
	return gorm.Delete(&App{Id: id})
}

//查询所有App
func QueryAllApp() []App {
	apps := new([]App)
	gorm.QueryAll(apps)
	return *apps
}
