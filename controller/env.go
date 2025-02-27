package controller

import (
    "os"
    "log"
	"github.com/joho/godotenv"
)

func GoDotEnvVariable(key string) string {
	err := godotenv.Load("local.env")
	if err != nil {
        log.Println(err)
		log.Fatalf("Loading .env file error")
	}

	return os.Getenv(key)

}
