package service

import (
	"gitchat/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 用户模块
// @Tags GetUserList
// @Success 200 {json} {code:200,message:{}]}
// @Router /user/get_user_list [get]
func GetUserList(c *gin.Context) {
	data := models.GetUserList()
	c.JSON(http.StatusOK, gin.H{
		"message": data,
	})
}

// 用户模块
// @Summary 添加用户
// @Description 用于添加新用户的接口
// @Tags AddUser
// @Accept json
// @Produce json
// @Param name query string false "用户名"
// @Param pass query string false "密码"
// @Param pass2 query string false "确认密码"
// @Success 200 {object} map[string]interface{} "成功响应"
// @Failure 400 {object} map[string]interface{} "请求参数错误"
// @Router /user/add_user [post]
func AddUser(c *gin.Context) {
	user := &models.UserBasic{
		Name:     c.Query("name"),
		PassWord: c.Query("pass"),
	}
	//验证用户名是否重复
	name := models.GetUserName(user.Name)
	Password2 := c.Query("pass2")
	if user.PassWord != Password2 {
		c.JSON(-1, gin.H{"message": "密码不一致"})
		return
	}
	if name == user.Name {
		c.JSON(-1, gin.H{"message": "已存在的用户名"})
		return
	}
	err := models.AddUser(user)
	if err != nil {
		c.JSON(-1, gin.H{"message": "添加失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "用户添加成功"})
}

// 用户模块
// @Summary 删除用户
// @Description 删除用户
// @Tags DeleteUser
// @Accept json
// @Produce json
// @Param id query int false "用户ID"
// @Router /user/delete_user [get]
func DeleteUser(c *gin.Context) {
	id := c.Query("id")
	int_id, _ := strconv.Atoi(id)
	err := models.DeleteUser(int_id)
	if err != nil {
		c.JSON(-1, gin.H{"message": "删除失败"})
		return
	}
	c.JSON(-1, gin.H{"message": "删除成功"})
}

// 用户模块
// @Summary 修改用户
// @Description 修改用户
// @Tags UpdateUser
// @Accept json
// @Produce json
// @Param id query string true "ID"
// @Param name query string true "用户名"
// @Param pwd query  string true "密码"
// @Router /user/edit [post]
func UpdateUser(c *gin.Context) {
	user := &models.UserBasic{
		Name:     c.PostForm("name"),
		PassWord: c.PostForm("pass"),
	}
	id := c.PostForm("id")
	user_id, _ := strconv.Atoi(id)
	user.ID = uint(user_id)
	err := models.UpdateUser(user)
	if err != nil {
		c.JSON(-1, gin.H{"message": "更新失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "更新成功"})
}
