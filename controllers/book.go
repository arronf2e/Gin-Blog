package controllers

import (
	"gin-login/apihelpers"
	"gin-login/models"
	"gin-login/services"

	"github.com/gin-gonic/gin"
)

// GetAllBook 获取所有书
func GetAllBook(c *gin.Context) {
	var book []models.Book
	err := services.GetAllBook(&book)
	if err != nil {
		apihelpers.RespondJSON(c, 404, book)
	} else {
		apihelpers.RespondJSON(c, 200, book)
	}
}

// AddNewBook 新增
func AddNewBook(c *gin.Context) {
	var book []models.Book
	c.BindJSON(&book)
	err := services.AddNewBook(&book)
	if err != nil {
		apihelpers.RespondJSON(c, 404, book)
	} else {
		apihelpers.RespondJSON(c, 200, book)
	}
}

// GetOneBook 根据id获取书
func GetOneBook(c *gin.Context) {
	id := c.Params.ByName("id")
	var book []models.Book
	err := services.GetOneBook(&book, id)
	if err != nil {
		apihelpers.RespondJSON(c, 404, book)
	} else {
		apihelpers.RespondJSON(c, 200, book)
	}
}

// PutOneBook 根据id修改书
func PutOneBook(c *gin.Context) {
	id := c.Params.ByName("id")
	var book []models.Book
	err := services.GetOneBook(&book, id)
	if err != nil {
		apihelpers.RespondJSON(c, 404, book)
	}
	c.BindJSON(&book)
	err = services.PutOneBook(&book, id)
	if err != nil {
		apihelpers.RespondJSON(c, 404, book)
	} else {
		apihelpers.RespondJSON(c, 200, book)
	}
}

// DeleteBook 删除书
func DeleteBook(c *gin.Context) {
	var book []models.Book
	id := c.Params.ByName("id")
	err := services.DeleteBook(&book, id)
	if err != nil {
		apihelpers.RespondJSON(c, 404, book)
	} else {
		apihelpers.RespondJSON(c, 200, book)
	}
}
