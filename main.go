package main

import (
	"gitchat/router"
	"gitchat/utils"
)

func main() {
	utils.ConfigInit()
	utils.InitMysql()
	r := router.Router()
	r.Run(":8000")
}
