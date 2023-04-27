package router

import (
	"golang-micro/controllers"

	"github.com/gin-gonic/gin"
)

func RouterApi() *gin.Engine {
	r := gin.Default()
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())
	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	v1 := r.Group("api")
	{
		v1.GET("data", controllers.Index)
		v1.GET("data/:id", controllers.Show)
		v1.POST("data", controllers.Create)
		v1.PUT("data/:id", controllers.Update)
		v1.DELETE("data/:id", controllers.Delete)
	}

	return r
}
