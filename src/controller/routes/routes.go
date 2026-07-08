package routes

import (
	"github.com/Maryszxxx/gocrud.git/src/controller"
	"github.com/Maryszxxx/gocrud.git/src/model"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup, userController controller.UserControllerInterface) {

	r.GET("/getUserById/:userId", model.VerifyTokenMiddleware, userController.FindUserByID)

	r.GET("/getUserByEmail/:email", model.VerifyTokenMiddleware, userController.FindUserByEmail)

	r.POST("/createUser", userController.CreateUser)

	r.PUT("/updateUser/:userId", userController.UpdateUser)

	r.DELETE("/deleteUser/:userId", userController.DeleteUser)

	r.POST("/login", userController.LoginUser)
}
