package utils

import "github.com/spf13/viper"

// 设置配置文件
func ConfigInit() {
	//配置文件名称
	viper.SetConfigName("app")
	//配置文件路径
	viper.AddConfigPath("ginchat/config")
	viper.ReadInConfig()

}
func InitMysql() {

}
