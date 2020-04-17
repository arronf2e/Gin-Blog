package services

import (
	"gin-blog/config"
	"gin-blog/models"
)

// GetAllPost 获取所有文章
func GetAllPost(p *[]models.Post) (err error) {
	if err = config.DB.Find(p).Error; err != nil {
		return err
	}
	return nil
}

// AddNewPost 添加新文章
func AddNewPost(p *[]models.Post) (err error) {
	if err = config.DB.Create(p).Error; err != nil {
		return err
	}
	return nil
}

// GetOnePost 根据 Postid 获取某一文章
func GetOnePost(p *[]models.Post, id string) (err error) {
	if err := config.DB.Where("id = ?", id).First(p).Error; err != nil {
		return err
	}
	return nil
}

// PutOnePost 修改文章信息
func PutOnePost(p *[]models.Post, id string) (err error) {
	config.DB.Where("id = ?", id).Save(p)
	return nil
}

// DeletePost 删除文章
func DeletePost(p *[]models.Post, id string) (err error) {
	config.DB.Where("id = ?", id).Delete(p)
	return nil
}
