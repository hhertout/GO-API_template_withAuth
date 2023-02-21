package router

import (
	"API_go/go_test/config"
	"API_go/go_test/controllers"
	"API_go/go_test/middleware"

	"github.com/gin-gonic/gin"
)

func Routes() {
	router := gin.New()

	config.CORSConfig(router)

	router.POST("/signup", controllers.Signup)
	router.POST("/login", controllers.Login)
	router.GET("/validate", middleware.RequireAuth, controllers.Validate)

	router.Run()
}
