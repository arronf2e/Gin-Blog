package models

import "github.com/jinzhu/gorm"

// Tag struct
type Tag struct {
	gorm.Model
	Name  string
	Color string
}

// TableName Tag
func (b *Tag) TableName() string {
	return "tag"
}
