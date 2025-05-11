package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening/:id", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "GET Oppening",
			})
		})
		v1.POST("/opening", func(ctx *gin.Context) {
			ctx.JSON(http.StatusCreated, gin.H{
				"message": "POST Oppening",
			})
		})
		v1.PUT("/opening/:id", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "PUT Oppening",
			})
		})
		v1.DELETE("/opening/:id", func(ctx *gin.Context) {
			ctx.JSON(http.StatusNoContent, gin.H{
				"message": "DELETE Opening",
			})
		})
		v1.GET("/openings", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "GET Oppening",
			})
		})
	}
}
