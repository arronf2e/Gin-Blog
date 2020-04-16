package main

import (
	"fmt"
	"gin-login/config"
	"gin-login/models"
	"gin-login/routers"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
)

var err error

func main() {
	config.DB, err = gorm.Open("mysql", "")
	if err != nil {
		fmt.Println("status", err)
	}
	defer config.DB.Close()
	config.DB.AutoMigrate(&models.Book{})
	routers.SetupRouter()
}