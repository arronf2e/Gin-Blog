package models

import "github.com/jinzhu/gorm"
// User declare
type User struct {
	gorm.Model
	Username		string
	Password		string
}

// TableName User
func (b *User) TableName() string {
	return "user"
}