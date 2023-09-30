package router

import "github.com/gin-gonic/gin"

// Initialize gin router and routes for the application
func Initialize() {
	// Initialize the router
	router := gin.Default()

	// Initialize the routes
	initializeRoutes(router)

	// Run the server on port 8080
	err := router.Run(":8080")
	if err != nil {
		println(err)
		return
	}
}
