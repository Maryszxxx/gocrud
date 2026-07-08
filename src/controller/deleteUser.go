package controller

import (
	"net/http"

	"github.com/Maryszxxx/gocrud.git/src/config/logger"
	"github.com/Maryszxxx/gocrud.git/src/config/rest_err"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.uber.org/zap"
)

// DeleteUser godoc
// @Summary Delete User
// @Description Deletes a user based on the ID provided as a parameter
// @Tags Users
// @Accept json
// @Produce json
// @Param userId path string true "ID of the user to be deleted"
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200 "OK"
// @Failure 400 {object} rest_err.RestErr
// @Failure 500 {object} rest_err.RestErr
// @Router /deleteUser/{userId} [delete]

func (uc *userControllerInterface) DeleteUser(c *gin.Context) {
	logger.Info("Init deleteUser controller",
		zap.String("journey", "deleteUser"),
	)

	userId := c.Param("userId")
	if _, err := bson.ObjectIDFromHex(userId); err != nil {
		errRest := rest_err.NewBadRequestError("Invalid userId, must be a hex value")
		c.JSON(int(errRest.Code), errRest)
		return
	}

	err := uc.service.DeleteUser(userId)
	if err != nil {
		logger.Error(
			"Error trying to call deleteUser service",
			err,
			zap.String("journey", "deleteUser"))
		c.JSON(int(err.Code), err)
		return
	}

	logger.Info(
		"deleteUser controller executed successfully",
		zap.String("userId", userId),
		zap.String("journey", "deleteUser"))

	c.Status(http.StatusOK)
}
