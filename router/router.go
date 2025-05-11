package router

import "github.com/gin-gonic/gin"

func Initialize() {
	
	// Initialize router
	router := gin.Default()

	// Initialize routes
	initializeRoutes(router)	


	// Run server
	router.Run(":8080")
}