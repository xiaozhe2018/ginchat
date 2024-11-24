package utils

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 设置配置文件
func ConfigInit() {
	//配置文件名称
	viper.SetConfigName("app")
	//配置文件路径
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("error : ", err)
	}
	// fmt.Println("config app:", viper.Get("mysql"))

}
func InitMysql() *gorm.DB {
	//获取配置文件
	dsn := viper.GetString("mysql.dns")
	DB, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// user := models.UserBasic{}
	// DB.Find(&user)
	// fmt.Println("data:", user)
	return DB
}
