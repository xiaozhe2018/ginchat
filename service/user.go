package service

import (
	"gitchat/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUserList
// @Tags GetUserList
// @Success 200 {json} {code:200,message:{}]}
// @Router /user/get_user_list [get]
func GetUserList(c *gin.Context) {
	data := models.GetUserList()
	c.JSON(http.StatusOK, gin.H{
		"message": data,
	})
}
