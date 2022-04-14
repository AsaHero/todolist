package handler

import "github.com/gin-gonic/gin"

type Handler struct {

}

func (h* Handler) InitRouter() *gin.Engine {
	router := gin.New()

	auth := router.Group("/")
	{
		auth.POST("/sign-in")
		auth.POST("/sign-up")
	}

	api := router.Group("/api")
	{
		lists := api.Group("/lists")
		{
			lists.POST("/")
			lists.GET("/")
			lists.GET("/:id")
			lists.PUT("/:id")
			lists.DELETE("/:id")

			items := lists.Group(":id/items")
			{
				items.POST("/")
				items.GET("/")
				items.GET("/:id_item")
				items.PUT("/:id_item")
				items.DELETE("/:id_item")
			}
		}
	}

	return router
}