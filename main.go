package main

import (
	"bicycle-rental/configs"
	"context"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	var err error
	environment := os.Getenv("ENV")
	if environment == "production" {
		err = godotenv.Load(".env")
	} else {
		err = godotenv.Load(".env")
	}
	if err != nil {
		panic(err)
	}
	fmt.Println("ENV: ", environment)

	client, err := configs.ConnectDb()
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(context.TODO())

	r := gin.Default()

	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	r.Run(port)
}
