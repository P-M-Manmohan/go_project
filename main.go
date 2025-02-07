package main

import (
	"log"
	"project/login/database"
	"project/login/routes"

	"github.com/gin-gonic/gin"
)


func main(){

    db := database.Connect()
    
     defer db.Close()

    router:= gin.Default()
    
    routes.SetUpRoutes(router, db)

    log.Println("Server running on port 8080")

    router.Run("localhost:8080")
}
