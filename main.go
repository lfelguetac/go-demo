package main

import (
	"fmt"
	"go-demo/config"
	"go-demo/delivery/http"
	"go-demo/utils"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	appEnv := os.Getenv("APP_ENV")
	if appEnv != "prod" {
		_err := godotenv.Load()
		if _err != nil {
			fmt.Println("Error loading .env file" + _err.Error())
		}
	}

	db := config.SetupDBConnection()
	defer config.CloseDBConnection(db)

	app := gin.Default()

	port := utils.GetStringEnv("APP_PORT", "8085")

	userController := GetDependencies(db)

	http.NewAppHandler(app, userController)

	log.Printf("Server stopped, err: %v", app.Run(":"+port))

}
