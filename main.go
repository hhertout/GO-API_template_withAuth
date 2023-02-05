package main

import (
	"API_go/go_test/initializers"
	"API_go/go_test/router"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDatabase()
	initializers.Migrate()
}

func main() {
	router.UserRoutes()
}
