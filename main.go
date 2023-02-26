package main

import (
	"bubble/dao"
	"bubble/models"
	"bubble/routers"
)

func main() {

	

	err := dao.InitMysql()
	if err != nil {
		panic(err)
	}
	dao.DB.AutoMigrate(&models.Todo{})
	// err := dao.DB.AutoMigrate(&models.Todo{})
	// if err != nil {
	// 	return
	// }
	//注册路由
	r := routers.SetupRouter()
	r.Run()
}
