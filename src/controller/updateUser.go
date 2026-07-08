package controller

import (
	"net/http"
	"strings"

	"github.com/Maryszxxx/gocrud.git/src/config/logger"
	"github.com/Maryszxxx/gocrud.git/src/config/rest_err"
	"github.com/Maryszxxx/gocrud.git/src/config/validation"
	"github.com/Maryszxxx/gocrud.git/src/controller/model/request"
	"github.com/Maryszxxx/gocrud.git/src/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.uber.org/zap"
)

// UpdateUser godoc
// @Summary Update User
// @Description Updates user details based on the ID provided as a parameter
// @Tags Users
// @Accept json
// @Produce json
// @Param userId path string true "ID of the user to be updated"
// @Param userRequest body request.UserUpdateRequest true "User information for update"
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200 "OK"
// @Failure 400 {object} rest_err.RestErr
// @Failure 500 {object} rest_err.RestErr
// @Router /updateUser/{userId} [put]

func (uc *userControllerInterface) UpdateUser(c *gin.Context) {
	userId := c.Param("userId")
	var userRequest request.UserUpdateRequest

	logger.Info("Init updateUser controller",
		zap.String("journey", "updateUser"),
	)

	if err := c.ShouldBindJSON(&userRequest); err != nil || strings.TrimSpace(userId) == "" {
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "updateUser"))
		errRest := validation.ValidateUserError(err)

		c.JSON(int(errRest.Code), errRest)
		return
	}

	if _, err := bson.ObjectIDFromHex(userId); err != nil {
		errRest := rest_err.NewBadRequestError("Invalid userId, must be a hex value")
		c.JSON(int(errRest.Code), errRest)
		return
	}

	domain := model.NewUserUpdateDomain(
		userRequest.Name,
		userRequest.Age,
	)

	err := uc.service.UpdateUser(userId, domain)
	if err != nil {
		logger.Error(
			"Error trying to call updateUser service",
			err,
			zap.String("journey", "updateUser"))
		c.JSON(int(err.Code), err)
		return
	}

	logger.Info(
		"updateUser controller executed successfully",
		zap.String("userId", userId),
		zap.String("journey", "updateUser"))

	c.Status(http.StatusOK)
}
