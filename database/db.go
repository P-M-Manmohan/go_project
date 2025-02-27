package database

import (
	"log"
    "fmt"

	"github.com/jmoiron/sqlx"
    "github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
    "project/login/controller"
)

func Connect() (*sqlx.DB) {
    dbName:=controller.GoDotEnvVariable("DB_NAME")
    dbUserName:=controller.GoDotEnvVariable("DB_USERNAME")
    dbPassword:=controller.GoDotEnvVariable("DB_PASSWORD")
    dbHost:=controller.GoDotEnvVariable("DB_HOST") 
    dbPort:=controller.GoDotEnvVariable("DB_PORT")

    connStr := fmt.Sprintf("user=%s dbname=%s sslmode=disable password=%s host=%s port=%s",dbUserName,dbName, dbPassword, dbHost, dbPort)

    db, err := sqlx.Connect("postgres", connStr)
     if err!= nil {
            log.Fatalln(err)
    }
    if err := db.Ping(); err !=nil{
        log.Fatal(err)
    }else{
        log.Println("Success")
    }

    runMigrations(db)
    return db

}

func runMigrations(db *sqlx.DB) {
    driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
    if err != nil {
        log.Fatal("Failed to create migration driver:", err)
    }

    m, err := migrate.NewWithDatabaseInstance(
        "file:///migrations", // Path to migration files
        "postgres",
        driver,
    )
    if err != nil {
        log.Fatal("Migration setup failed:", err)
    }

    if err := m.Up(); err != nil && err != migrate.ErrNoChange {
        log.Fatal("Error applying migrations:", err)
    }

    log.Println("Migrations applied successfully")
}

