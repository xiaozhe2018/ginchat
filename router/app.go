package router

import (
	"gitchat/docs"
	"gitchat/service"

	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	docs.SwaggerInfo.BasePath = ""
	//swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	//首页
	r.GET("index", service.GetIndex)
	//获取用户列表
	r.GET("user/get_user_list", service.GetUserList)
	//添加用户
	r.POST("user/add_user", service.AddUser)
	//删除用户
	r.GET("user/delete_user", service.DeleteUser)
	//修改用户
	r.POST("user/edit", service.UpdateUser)
	//用户登录
	r.POST("user/login", service.LoginUser)
	return r
}
