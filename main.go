package main

import (
	"context"
	"log"

	"github.com/Maryszxxx/gocrud.git/src/config/database/mongodb"
	"github.com/Maryszxxx/gocrud.git/src/config/logger"
	"github.com/Maryszxxx/gocrud.git/src/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// @title Meu Primeiro CRUD em Go
// @version 1.0
// @description API for crud operations on users
// @host localhost:8080
// @BasePath /
func main() {
	logger.Info("About to start user application")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		log.Fatalf(
			"Error trying to connect to database, error=%s \n",
			err.Error())
		return
	}

	userController := initDependencies(database)

	router := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

// docker container run -d -p 27017:27017 mongo (inicializador mongodb)
