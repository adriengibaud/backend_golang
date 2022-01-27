package main

import (
	"backend_golang/configs"
	"backend_golang/middleware"
	"backend_golang/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

		  config := cors.DefaultConfig()
  config.AllowOrigins = []string{"http://localhost:8080"}

		router.Use(cors.New(config))

		router.Use(middleware.BasicMiddleware())
    
    //run database
    configs.ConnectDB()

		routes.UserRoute(router)

    router.Run("localhost:7789")


}