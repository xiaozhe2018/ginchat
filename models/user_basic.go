package models

import (
	"gitchat/utils"

	"gorm.io/gorm"
)

// 结构体
type UserBasic struct {
	gorm.Model
	Name          string
	PassWord      string
	Phone         string
	Email         string
	Identity      string //用户唯一标识
	ClentIp       string
	ClientPort    string
	LoginTime     uint64
	HeartbeatTime uint64
	LoginOutTime  uint64
	IsLogout      bool
	DeviceInfo    string
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}

func GetUserList() []*UserBasic {
	data := make([]*UserBasic, 10)
	utils.DB.Find(&data)
	return data
}

// 新增
func AddUser(user *UserBasic) (err error) {
	err = utils.DB.Create(user).Error
	return err
}

// 修改
func UpdateUser(user *UserBasic) (err error) {
	err = utils.DB.Save(user).Error
	return err
}

//删除

//查询
