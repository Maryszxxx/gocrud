package main

import (
	"log"

	"github.com/Maryszxxx/gocrud.git/src/controller"
	"github.com/Maryszxxx/gocrud.git/src/controller/model/service"
	"github.com/Maryszxxx/gocrud.git/src/controller/model/view/test/config/logger"
	"github.com/Maryszxxx/gocrud.git/src/controller/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	logger.Info("Starting the application...")

	//init dependencies
	service := service.NewUserDomainService()
	userController := controller.NewUserControllerInterface(service)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
