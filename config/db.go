package config

import (
	"fmt"
	"gin-blog/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB 定义
var DB *gorm.DB

// InitSQL 初始化数据库
func InitSQL() {
	var err error
	DB, err = gorm.Open("mysql", "root:root@tcp(127.0.0.1:32771)/ginblog?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("status", err)
	}
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Post{})
	DB.AutoMigrate(&models.Tag{})
}

// CloseSQL 关闭数据库连接
func CloseSQL() {
	defer DB.Close()
}
