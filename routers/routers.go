package routers

import (
	"gin-login/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter()  {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("book", controllers.GetAllBook)
		v1.POST("book", controllers.AddNewBook)
		v1.GET("book/:id", controllers.GetOneBook)
		v1.PUT("book/:id", controllers.PutOneBook)
		v1.DELETE("book/:id", controllers.DeleteBook)
	}
	r.Run()
	//return v1
}