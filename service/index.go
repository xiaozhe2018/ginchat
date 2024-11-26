package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1
// GetIndex
// @Tags 首页
// @Success 200 {string} this is home
// @Router /index [get]
/**
 * @BasePath /api/v1
 * GetIndex
 * @param name query string true "name"
 * @Tags 首页
 * @Success 200 {string} this is home
 * Router /index [get]
 */
func GetIndex(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "this is home !!",
	})
}
