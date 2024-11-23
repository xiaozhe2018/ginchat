package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 处理index
func GetIndex(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "你好啊！！",
	})
}
