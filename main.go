package main

import "bubble/routers"

func main() {

	//注册路由
	r := routers.SetupRouter()
	r.GET("/", routers.IndexHandlers)
	r.Run()
}
