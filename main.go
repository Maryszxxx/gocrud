package main

import (
	"log"

	"github.com/Maryszxxx/gocrud.git/src/controller/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	logger.Info("Starting the application...")
	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup)
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
