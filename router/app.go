package router

import (
	"gitchat/service"

	"github.com/gin-gonic/gin"
)

// func Router() {
// 	r := gin.Default()
// 	r.GET("/ping", func(c *gin.Context) {
// 		c.JSON(http.StatusOK, gin.H{
// 			"message": "pong",
// 		})
// 	})
// 	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
// }

func Router() *gin.Engine {

	r := gin.Default()
	//首页
	r.GET("index", service.GetIndex)
	//获取用户列表
	r.GET("user/get_user_list", service.GetUserList)

	return r
}
