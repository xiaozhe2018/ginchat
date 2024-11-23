package main

import (
	"fmt"
	"gitchat/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:123456@tcp(127.0.0.1:23306)/ginchat?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	User := &models.UserBasic{}
	User.Name = "小哲"
	// Migrate the schema
	db.AutoMigrate(User)

	// Create
	db.Create(User)

	// Read
	// var product Product
	fmt.Println(db.First(User, 1)) // find product with integer primary key
	// db.First(User, "code = ?", "D42") // find product with code D42

	// Update - update product's price to 200
	db.Model(User).Update("PassWord", "123")
	// // Update - update multiple fields
	// db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	// db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// // Delete - delete product
	// db.Delete(&product, 1)
}
