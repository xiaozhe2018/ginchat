package service

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"gitchat/jwt"
	"gitchat/models"
	"gitchat/utils"
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
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
// @Param phone query string false "手机号"
// @Param email query string false "邮箱"
// @Success 200 {object} map[string]interface{} "成功响应"
// @Failure 400 {object} map[string]interface{} "请求参数错误"
// @Router /user/add_user [post]
func AddUser(c *gin.Context) {
	user := &models.UserBasic{
		Name:     c.Query("name"),
		PassWord: c.Query("pass"),
		Phone:    c.Query("phone"),
		Email:    c.Query("email"),
	}
	msgMap := map[string]string{
		"name":  "用户名已存在",
		"phone": "手机号已存在",
		"email": "邮箱已存在",
	}
	//验证用户名是否重复
	// name := models.GetUserName(user.Name)
	Password2 := c.Query("pass2")
	if user.PassWord != Password2 {
		c.JSON(-1, gin.H{"message": "密码不一致"})
		return
	}
	//验证字段是否存在
	name, d := models.GetUserField("name", user.Name)
	if d != "" {
		fmt.Println("d:", d)
		c.JSON(-1, gin.H{"message": msgMap[name]})
		return
	}
	phone, d := models.GetUserField("phone", user.Phone)
	if d != "" {
		fmt.Println("d:", d)
		c.JSON(-1, gin.H{"message": msgMap[phone]})
		return
	}
	email, d := models.GetUserField("email", user.Email)
	if d != "" {
		fmt.Println("d:", d)
		c.JSON(-1, gin.H{"message": msgMap[email]})
		return
	}
	// if name == user.Name {
	// 	c.JSON(-1, gin.H{"message": "已存在的用户名"})
	// 	return
	// }
	//生成一个6位的随机数
	randomBytes := make([]byte, 6)
	_, err2 := rand.Read(randomBytes)
	if err2 != nil {
		panic(err2)
	}
	salt := base64.URLEncoding.EncodeToString(randomBytes)
	//CreatePass创建加密密码
	user.PassWord = utils.CreatePass(user.PassWord, salt)
	user.Salt = salt
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
// @Param email query  string true "邮箱"
// @Param phone query  string true "电话号码"
// @Router /user/edit [post]
func UpdateUser(c *gin.Context) {
	user := &models.UserBasic{
		Name:     c.PostForm("name"),
		PassWord: c.PostForm("pass"),
	}
	id := c.PostForm("id")
	user_id, _ := strconv.Atoi(id)
	user.ID = uint(user_id)
	user.Email = c.PostForm("email")
	user.Phone = c.PostForm("phone")
	//另外一种写法要在&models.UserBasic上配置
	// _, vlidate_err := govalidator.ValidateStruct(user)
	// if vlidate_err != nil {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"code": -1,
	// 		"msg":  "数据验证错误",
	// 	})
	// 	return
	// }
	res := govalidator.IsEmail(user.Email)
	if res != true {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "邮箱格式不正确",
		})
		return
	}
	phoneRegex := "^1[3-9]\\d{9}$"
	PhoneRes := govalidator.Matches(user.Phone, phoneRegex)
	if PhoneRes != true {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "手机号格式不正确",
		})
		return
	}
	// if govalidator.IsP
	err := models.UpdateUser(user)
	if err != nil {
		c.JSON(-1, gin.H{"message": "更新失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "更新成功"})
}

// 用户模块
// @Summary 用户登录
// @Description 用户登录
// @Tags LoginUser
// @Accept json
// @Produce json
// @Param name query string true "用户名/邮箱/手机号码"
// @Param pass query string true "密码"
// @Router /user/login [post]
func LoginUser(c *gin.Context) {
	name := c.PostForm("name")
	pass := c.PostForm("pass")
	row := models.GetUserAllField(name)
	if row == nil {
		// c.JSON(-1, gin.H{"message": "用户不存在"})
		utils.ErrorJson(c, -1, "用户不存在")
		return
	}
	//验证密码
	validPassword := utils.ValidPassword(pass, row.Salt, row.PassWord)
	if !validPassword {
		// c.JSON(-1, gin.H{"message": "密码错误"})
		utils.ErrorJson(c, -2, "密码错误")
		return
	}
	//发放token
	//构造jwt
	token := jwt.GenToken(string(row.ID), row.Name)
	// c.JSON(http.StatusOK, gin.H{"message": "登录成功", "user": row})
	//简单token的处理先实现功能，其实可以研究下gin/jwt
	row.Identity = token
	utils.Success(200, row, "success", c)
}
