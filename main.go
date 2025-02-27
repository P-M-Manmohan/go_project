package main

import (
	"log"
	"project/login/database"
	"project/login/routes"

	"github.com/gin-gonic/gin"
)


func main(){

    //for{
    //    print("hello")
    //}
    db := database.Connect()
    
     defer db.Close()

    router:= gin.Default()
    router.Use(gin.Logger())
    routes.SetUpRoutes(router, db)

    log.Println("Server running on port 8080")

    router.Run("0.0.0.0:8080")// for docker 0.0.0.0:8080
}
