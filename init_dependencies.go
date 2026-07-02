package main

import (
	"github.com/Maryszxxx/gocrud.git/src/controller"
	"github.com/Maryszxxx/gocrud.git/src/model/repository"
	"github.com/Maryszxxx/gocrud.git/src/model/service"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func initDependencies(
	database *mongo.Database,
) controller.UserControllerInterface {
	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	return controller.NewUserControllerInterface(service)
}
