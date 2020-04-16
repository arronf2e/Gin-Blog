package models

import "github.com/jinzhu/gorm"

// Book 定义 book 数据模型
type Book struct {
	gorm.Model
	Name     string `json:"name"`
	Author   string `json:"author"`
	Category string `json:"category"`
}

// TableName Book
func (b *Book) TableName() string {
	return "book"
}
