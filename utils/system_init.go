package utils

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// 设置配置文件
var DB *gorm.DB

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
func InitMysql() {
	//自定义日志
	new_log := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, //慢查询阈值
			LogLevel:      logger.Info, //级别
			Colorful:      true,        //开启颜色
		},
	)
	//获取配置文件
	dsn := viper.GetString("mysql.dns")
	DB, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: new_log})
}
