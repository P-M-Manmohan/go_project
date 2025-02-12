package middleware

import (
	"bytes"
	"log"
	"net/http"
	"project/login/controller"
	"project/login/model"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jmoiron/sqlx"
)


type responseRecorder struct {
    gin.ResponseWriter
    body *bytes.Buffer
    statusCode int
}

func (r *responseRecorder) WriteHeader(code int) {
	r.statusCode = code                
	r.ResponseWriter.WriteHeader(code)
}

func (r *responseRecorder) Write(b []byte) (int, error) {
    r.body.Write(b)               
    return len(b), nil            
}

type AuthMiddleware struct {
	DB *sqlx.DB
}

func NewAuthMiddleware(db *sqlx.DB) *AuthMiddleware {
	return &AuthMiddleware{DB: db}
}


func (uc *AuthMiddleware) TokenAuth(c *gin.Context){
    var input model.User

    if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input in middleware"})
		c.Abort()
        return
	}
    tokenString := c.GetHeader("Token")
    
    secretKey := []byte(controller.GoDotEnvVariable("SECRET_KEY"))
    err:= uc.DB.QueryRow("SELECT token,role FROM users WHERE name = $1", input.Name).
    Scan(&input.JWT,&input.Role)
    if err!=nil{
        log.Println(err)
        c.JSON(http.StatusForbidden,gin.H{"error":"Access Denied middleware"})
        c.Abort()
        return
    }
    

    if input.JWT != tokenString{
        c.JSON(http.StatusForbidden, gin.H{"error":"Access Denied middleware"})
        c.Abort()
        return
    }

    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
      return secretKey, nil
    })
  
    if err != nil {
        c.JSON(http.StatusBadGateway, gin.H{"error":"Error during Authentication middleware"})
        c.Abort()
        return
    }

  
    if !token.Valid {
        c.JSON(http.StatusForbidden, gin.H{"error":"Access Denied middleware"})
        c.Abort()
        return
    }
    log.Println("Authenticated")
    c.Set("authenticatedUser", input)
    c.Next()

}

