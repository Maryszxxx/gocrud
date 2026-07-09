package routes

import (
	_ "github.com/Maryszxxx/gocrud.git/docs"
	"github.com/Maryszxxx/gocrud.git/src/controller"
	"github.com/Maryszxxx/gocrud.git/src/model"
	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRoutes(r *gin.RouterGroup, userController controller.UserControllerInterface) {

	r.GET("/getUserById/:userId", model.VerifyTokenMiddleware, userController.FindUserByID)

	r.GET("/getUserByEmail/:email", model.VerifyTokenMiddleware, userController.FindUserByEmail)

	r.POST("/createUser", userController.CreateUser)

	r.PUT("/updateUser/:userId", userController.UpdateUser)

	r.DELETE("/deleteUser/:userId", userController.DeleteUser)

	r.POST("/login", userController.LoginUser)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	//http://localhost:8080/swagger/index.html#/
}
