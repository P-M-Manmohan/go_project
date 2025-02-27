package routes

import (
	"log"
	"project/login/controller"
	"project/login/middleware"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func SetUpRoutes(router *gin.Engine, db *sqlx.DB) {
    log.Println("router")
	userControllers := controller.NewUserController(db)
	router.GET("/", userControllers.Home)
	router.POST("/login", userControllers.Login)
	router.POST("/signin", userControllers.CreateUser)
	api := router.Group("/auth")
	api.Use(middleware.NewAuthMiddleware(db).TokenAuth)
	{
		api.GET("/users", userControllers.GetUsers)
		api.DELETE("/deleteuser", userControllers.DeleteUser)
	}

}
