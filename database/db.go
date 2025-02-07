package database

import (
	"log"
	"os"
    "fmt"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
    //"project/login/controller"
)

func GoDotEnvVariable(key string) string{
    err := godotenv.Load(".env")
    if err!=nil{
        log.Fatalf("Error loading .env file")
    }

    return os.Getenv(key)
    
}

func Connect() (*sqlx.DB) {
    dbName:=GoDotEnvVariable("DB_NAME")
    dbUserName:=GoDotEnvVariable("DB_USERNAME")
    dbPassword:=GoDotEnvVariable("DB_PASSWORD")
    dbHost:=GoDotEnvVariable("DB_HOST")

    connStr := fmt.Sprintf("user=%s dbname=%s sslmode=disable password=%s host=%s",dbUserName,dbName, dbPassword, dbHost)

    db, err := sqlx.Connect("postgres", connStr)
     if err!= nil {
            log.Fatalln(err)
    }
    if err := db.Ping(); err !=nil{
        log.Fatal(err)
    }else{
        log.Println("Success")
    }
    return db

}
