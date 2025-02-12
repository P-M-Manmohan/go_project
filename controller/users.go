package controller

import (
	"log"
	"net/http"
	"project/login/model"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type UserController struct {
	DB *sqlx.DB
}

func NewUserController(db *sqlx.DB) *UserController {
	return &UserController{DB: db}
}

func (uc *UserController) GetUsers(c *gin.Context) {
	var users []model.User
    var input model.User
    authUser, exists := c.Get("authenticatedUser")
    if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
    input, ok := authUser.(model.User)
    if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Type assertion failed"})
		return
	}


    log.Println(input)
    if input.Role=="admin"{
        err := uc.DB.Select(&users, "SELECT * from users")
	    if err != nil {
		    log.Println(err)
		    c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		    return
	    }
    }else{
        err := uc.DB.Select(&users, "SELECT * from users WHERE role='user'")
	    if err != nil {
		    log.Println(err)
		    c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		    return
	    }       
    }
	c.JSON(http.StatusOK, users)
}

func (uc *UserController) CreateUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

    salt, hashedPassword, err := HashPassword(user.Password)

    if err!= nil{
        c.JSON(http.StatusInternalServerError,gin.H{"error" : "Failed to hash password"})
    }
    user.Password=hashedPassword
    user.Salt=salt

	_, err = uc.DB.Exec("INSERT INTO users (name, email, password, salt) VALUES ($1, $2, $3, $4)", user.Name, user.Email, hashedPassword, salt)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
		return
	}
	c.JSON(http.StatusCreated, user)
}

func (uc *UserController) Login(c *gin.Context){
    
    var user, input model.User
    
    if err := c.ShouldBindJSON(&input); err!=nil{
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }
    
    err :=uc.DB.QueryRow("SELECT name,password,salt,token FROM users WHERE name=$1", input.Name).
    Scan(&user.Name,&user.Password,&user.Salt,&user.JWT)
    
    if err != nil {
        log.Println(err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed retrieving user data"})
        return
    }

    authorized := VerifyPassword(input.Password, user.Password, user.Salt)
    
    if authorized != true{
        c.JSON(http.StatusForbidden,"Incorrect Username and password")
        return
    }
    
    
    if user.JWT=="1"{
        user.JWT,err = GenerateTocken(input.Name)
        if err!=nil{
            log.Println(err)
            c.JSON(http.StatusInternalServerError, "Error Creating JWT")
        }
    }else{
        err=VerifyToken(user.JWT)
        if err!=nil{
            user.JWT,err = GenerateTocken(input.Name)
            if err!=nil{
                log.Println(err)
                c.JSON(http.StatusInternalServerError, "Error Creating JWT")
            }   
        }
    }

    _, err = uc.DB.Exec("UPDATE users SET token=$1 WHERE name=$2",user.JWT,input.Name)

    

    c.JSON(http.StatusAccepted,gin.H{ "Token" :user.JWT})
            log.Println(string(responseBody),"in response")

}

func (uc *UserController) DeleteUser(c *gin.Context){
    var user,input model.User
    if err := c.ShouldBindJSON(&input); err!=nil{
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }
    
    log.Println(input)
    err :=uc.DB.QueryRow("SELECT name,password,salt FROM users WHERE name=$1", input.Name).
    Scan(&user.Name,&user.Password,&user.Salt)
    if err != nil {
        log.Println(err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed retrieving user data"})
        return
    }

    authorized := VerifyPassword(input.Password, user.Password, user.Salt)
    
    if authorized != true{
        c.JSON(http.StatusForbidden,"Incorrect Username and password")
        return
    }

    _, err = uc.DB.Exec("DELETE FROM users WHERE name=$1", input.Name)

    c.JSON(http.StatusOK,"User Deleted")
}
