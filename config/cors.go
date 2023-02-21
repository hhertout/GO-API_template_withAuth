package config

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORSConfig(router *gin.Engine) {
	router.SetTrustedProxies([]string{"*"})
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{os.Getenv("DEST")}
	config.AllowCredentials = true
	router.Use(cors.New(config))
}
