package models

import (
	"gitchat/utils"
	"log"
	"time"

	"gorm.io/gorm"
)

// 结构体
type UserBasic struct {
	gorm.Model
	Name          string
	PassWord      string
	Phone         string `valid:"matches(^1[3-9]{1}\\d{9}$)"`
	Email         string `valid:"IsEmail"` //要下载govalidate库  go get github.com/asaskevich/govalidator
	Identity      string //用户唯一标识
	ClentIp       string
	Salt          string //随机数加盐
	ClientPort    string
	LoginTime     uint64
	HeartbeatTime uint64
	LoginOutTime  uint64
	IsLogout      bool
	DeviceInfo    string
}
type Model struct {
	ID        uint           `gorm:"primarykey"`
	CreatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}

func GetUserList() []*UserBasic {
	data := make([]*UserBasic, 10)
	utils.DB.Find(&data)
	return data
}

// 获取用户名
func GetUserName(name string) string {
	data := []*UserBasic{}
	err := utils.DB.Where("name=?", name).First(&data).Error
	if err != nil {
		log.Println("查询用户失败", err)
		return ""
	}
	return data[0].Name
}

// 返回用户全部字段
func GetUserAllField(name string) *UserBasic {
	data := &UserBasic{}
	utils.DB.Where("name=?", name).First(&data)
	return data
}

// 返回指定用户字段的信息
func GetUserField(field string, value string) (string, string) {
	data := []*UserBasic{}
	utils.DB.Where(field+"=?", value).Find(&data)
	if len(data) > 0 {
		return field, data[0].Name
	}
	return field, ""
}

// 新增
func AddUser(user *UserBasic) (err error) {
	err = utils.DB.Create(user).Error
	return err
}

// 修改
func UpdateUser(user *UserBasic) (err error) {
	// err = utils.DB.Save(user).Error
	err = utils.DB.Model(user).Where("id=?", user.ID).Updates(user).Error
	return err
}

// 删除
func DeleteUser(id int) (err error) {
	user := &UserBasic{}
	err = utils.DB.Delete(&user, id).Error
	return err
}

//查询
