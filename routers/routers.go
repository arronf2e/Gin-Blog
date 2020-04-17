package routers

import (
	"gin-blog/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("post", controllers.GetAllPost)
		v1.POST("post", controllers.AddNewPost)
		v1.GET("post/:id", controllers.GetOnePost)
		v1.PUT("post/:id", controllers.PutOnePost)
		v1.DELETE("post/:id", controllers.DeletePost)
	}
	r.Run()
	//return v1
}
