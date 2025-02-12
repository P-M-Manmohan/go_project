package routes

import (
	"project/login/controller"
	"project/login/middleware"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func SetUpRoutes(router *gin.Engine, db *sqlx.DB){
    userControllers := controller.NewUserController(db)
    router.POST("/login",userControllers.Login)
    api:= router.Group("/auth")
    api.Use(middleware.NewAuthMiddleware(db).TokenAuth)
    {
        api.GET("/users",userControllers.GetUsers)
        api.POST("/signin",userControllers.CreateUser)
        api.DELETE("/deleteuser",userControllers.DeleteUser)
    }

}
