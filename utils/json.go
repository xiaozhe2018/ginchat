package utils

import "github.com/gin-gonic/gin"

func Success(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(code, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}

func ErrorJson(c *gin.Context, code int, msg string) {
	c.JSON(code, gin.H{
		"code": code,
		"msg":  msg,
	})
	return
}
