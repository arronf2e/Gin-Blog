package controllers

import (
	"fmt"
	"gin-blog/apihelpers"
	"gin-blog/models"
	"gin-blog/services"

	"github.com/gin-gonic/gin"
)

// GetAllPost 获取所有文章
func GetAllPost(c *gin.Context) {
	var post []models.Post
	err := services.GetAllPost(&post)
	fmt.Println(err)
	if err != nil {
		apihelpers.RespondJSON(c, 500, err)
	} else {
		apihelpers.RespondJSON(c, 200, post)
	}
}

// AddNewPost 新增文章
func AddNewPost(c *gin.Context) {
	var post []models.Post
	c.BindJSON(&post)
	err := services.AddNewPost(&post)
	if err != nil {
		apihelpers.RespondJSON(c, 404, post)
	} else {
		apihelpers.RespondJSON(c, 200, post)
	}
}

// GetOnePost 根据id获取文章
func GetOnePost(c *gin.Context) {
	id := c.Params.ByName("id")
	var post []models.Post
	err := services.GetOnePost(&post, id)
	if err != nil {
		apihelpers.RespondJSON(c, 404, post)
	} else {
		apihelpers.RespondJSON(c, 200, post)
	}
}

// PutOnePost 根据id修改文章
func PutOnePost(c *gin.Context) {
	id := c.Params.ByName("id")
	var post []models.Post
	err := services.GetOnePost(&post, id)
	if err != nil {
		apihelpers.RespondJSON(c, 404, post)
	}
	c.BindJSON(&post)
	err = services.PutOnePost(&post, id)
	if err != nil {
		apihelpers.RespondJSON(c, 404, post)
	} else {
		apihelpers.RespondJSON(c, 200, post)
	}
}

// DeletePost 删除文章
func DeletePost(c *gin.Context) {
	var post []models.Post
	id := c.Params.ByName("id")
	err := services.DeletePost(&post, id)
	if err != nil {
		apihelpers.RespondJSON(c, 404, post)
	} else {
		apihelpers.RespondJSON(c, 200, post)
	}
}
