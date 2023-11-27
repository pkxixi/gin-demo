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

	global.DB = initial.Mysql()
	initial.RegisterTables()
	fmt.Println(global.DB)

	global.Redis = initial.Redis()
	router := routers.InitRouter()
	//router.Use(middleware.LoggerToFile("test.log"))
	addr := global.Config.System.Addr()
	global.Logger.Infof("server is running at: %s", addr)
	err := router.Run(addr)
	if err != nil {
		global.Logger.Fatalf("can not start system: %v", err)
		//log.Fatalf("can not start system: %v", err)
	}

}

//docker run -itd --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 mysql:5.7
