package service

import (
	"gitchat/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

/**
* 获取用户列表
 */
func GetUserList(c *gin.Context) {
	data := models.GetUserList()
	c.JSON(http.StatusOK, gin.H{
		"message": data,
	})
}
