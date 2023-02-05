package router

import (
	"API_go/go_test/controllers"
	"API_go/go_test/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes() {
	route := gin.Default()

	route.POST("/signup", controllers.Signup)
	route.POST("/login", controllers.Login)
	route.GET("/validate", middleware.RequireAuth, controllers.Validate)

	route.Run()
}
