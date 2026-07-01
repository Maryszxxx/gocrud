package controller

import (
	"fmt"
	"net/http"

	"github.com/Maryszxxx/gocrud.git/src/controller/model/view/test/config/logger"
	"github.com/Maryszxxx/gocrud.git/src/controller/model/view/test/config/rest_err/request"
	"github.com/Maryszxxx/gocrud.git/src/controller/model/view/test/config/rest_err/response"
	"github.com/Maryszxxx/gocrud.git/src/controller/model/view/test/config/validation"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap/zapcore"
)

func CreateUser(c *gin.Context) {
	logger.Info("CreateUser called",
		zapcore.Field{
			Key:    "journey",
			String: "CreateUser",
		},
	)
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error binding JSON: ", err)
		errRest := validation.ValidateUserError(err)

		c.JSON(http.StatusBadRequest, errRest)
		return
	}

	fmt.Println(userRequest)
	response := response.UserResponse{
		ID:    "1",
		Name:  userRequest.Name,
		Email: userRequest.Email,
		Age:   userRequest.Age,
	}
	c.JSON(http.StatusOK, response)
}
