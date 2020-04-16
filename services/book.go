package services

import (
	"gin-login/config"
	"gin-login/models"
)

// GetAllBook 获取所有书
func GetAllBook(b *[]models.Book) (err error) {
	if err = config.DB.Find(b).Error; err != nil {
		return err
	}
	return nil
}

// AddNewBook 添加新书
func AddNewBook(b *[]models.Book) (err error) {
	if err = config.DB.Create(b).Error; err != nil {
		return err
	}
	return nil
}

// GetOneBook 根据 bookid 获取某一本书
func GetOneBook(b *[]models.Book, id string) (err error) {
	if err := config.DB.Where("id = ?", id).First(b).Error; err != nil {
		return err
	}
	return nil
}

// PutOneBook 修改书本信息
func PutOneBook(b *[]models.Book, id string) (err error) {
	config.DB.Where("id = ?", id).Save(b)
	return nil
}

// DeleteBook 删除书本
func DeleteBook(b *[]models.Book, id string) (err error) {
	config.DB.Where("id = ?", id).Delete(b)
	return nil
}
