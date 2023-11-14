package main

import (
	"fmt"
	"go-blog/global"
	"go-blog/initial"
	"go-blog/routers"
)

func main() {
	initial.InitConf()
	fmt.Println(global.Config)

	global.Logger = initial.InitLogger()

	global.DB = initial.MysqlConnect()
	fmt.Println(global.DB)

	router := routers.InitRouter()
	addr := global.Config.System.Addr()
	global.Logger.Infof("server is running at: %s", addr)
	err := router.Run(addr)
	if err != nil {
		global.Logger.Fatalf("can not start system: %v", err)
	}

}
