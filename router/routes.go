package router

import (
	"github.com/pierrialexandervidmar/go-jobs/handler"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening/:id", handler.GetOpeningHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.PUT("/opening/:id", handler.UpdateOpeningHandler)
		v1.DELETE("/opening/:id", handler.DeleteOpeningHandler)
		v1.GET("/openings", handler.ListOpeningsHandler)
	}
}
