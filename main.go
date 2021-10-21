package main

import (
	"go.uber.org/zap"
	"log"
	"test/dao"
	"test/logger"
	"test/routers"
	"test/setting"
)

func main() {
	//初始化配置信息
	err := setting.Init()
	if err != nil{
		//不成功就没有然后了
		log.Printf("%#v",err)
		return
	}

	err = logger.Init(setting.Conf.Logconf)
	if err  != nil {
		log.Printf("%#v",err)
		return
	}
	defer zap.L().Sync()

	//初始化数据库
	err = dao.Initmysql(setting.Conf.Mysqlconf)
	if err != nil {
		log.Printf("%#v",err)
		zap.L().Error("mysql init err",zap.Error(err))
		return
	}
	defer dao.Close()

	r := routers.Setuprouter()
	r.Run(":8080")
}
