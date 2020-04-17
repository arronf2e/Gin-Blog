package models

import "github.com/jinzhu/gorm"

// Post declare
type Post struct {
	gorm.Model
	Title   string
	content string
	// Tags     []int
	// comments []interface{}
	author int
}

// TableName Post
func (b *Post) TableName() string {
	return "post"
}
