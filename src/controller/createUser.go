package controller

import (
	"net/http"

	"github.com/Maryszxxx/gocrud.git/src/controller/model"
	"github.com/Maryszxxx/gocrud.git/src/controller/model/view/test/config/logger"
	"github.com/Maryszxxx/gocrud.git/src/controller/model/view/test/config/rest_err/request"
	"github.com/Maryszxxx/gocrud.git/src/controller/model/view/test/config/validation"
	"github.com/Maryszxxx/gocrud.git/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func (uc *userControllerInterface) CreateUser(c *gin.Context) {

	logger.Info("CreateUser called",
		zap.String("journey", "CreateUser"),
	)

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {

		logger.Error("Error binding JSON: ", err)

		errRest := validation.ValidateUserError(err)

		c.JSON(http.StatusBadRequest, errRest)
		return
	}

	domain := model.NewUserDomain(
		userRequest.Name,
		userRequest.Password,
		userRequest.Email,
		userRequest.Age,
	)

	if err := uc.service.CreateUser(domain); err != nil {
		c.JSON(int(err.Code), err)
		return
	}

	logger.Info("User created successfully",
		zap.String("journey", "CreateUser"))

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domain))
}
