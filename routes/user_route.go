package routes

import (
	"backend_golang/controllers"
	"fmt"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine)  {
    fmt.Print(router)
    router.POST("/user", controllers.CreateUser()) //add this

    router.GET("/users", controllers.GetAllUsers())

    router.GET("/user/:userId", controllers.GetAUser())

    router.PUT("/user/:userId", controllers.EditAUser())

    router.DELETE("/user/:userId", controllers.DeleteAUser())
}