package routes

import (
	"project/login/controller"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func SetUpRoutes(router *gin.Engine, db *sqlx.DB){
    userControllers := controller.NewUserController(db)

    api:= router.Group("/api")
    {
        api.GET("/users",userControllers.GetUsers)
        api.POST("/signin",userControllers.CreateUser)
        api.POST("/login",userControllers.Login)
        api.DELETE("/deleteuser",userControllers.DeleteUser)
    }
}
