package main

import (
	"gitchat/router"
	"gitchat/utils"
)

func main() {
	//初始化配置文件
	utils.ConfigInit()
	//初始化mysql
	utils.InitMysql()
	//引入路由
	r := router.Router()
	//定义路由端口
	r.Run(":8000")
}
