package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/kalimoldayev02/api-golang/pkg/service"
)

type Handler struct {
	sevices *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{sevices: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-in", h.signIn)
		auth.POST("/sign-up", h.signUp)
	}

	// api := router.Group("/api")
	// {
	// 	lists := api.Group("/lists")
	// 	{
	// 		lists.POST("/")
	// 		lists.GET("/")
	// 		lists.GET("/:id")
	// 		lists.PUT("/:id")
	// 		lists.DELETE("/:id")

	// 		items := api.Group("/:id/items")
	// 		{
	// 			items.POST("/")
	// 			items.GET("/")
	// 			items.GET("/:item_id")
	// 			items.PUT("/:item_id")
	// 			items.DELETE("/:item_id")
	// 		}
	// 	}
	// }
	return router
}
