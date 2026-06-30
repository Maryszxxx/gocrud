package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Maryszxxx/gocrud.git/src/controller/model/view/test/config/rest_err/request"
	"github.com/Maryszxxx/gocrud.git/src/controller/model/view/test/config/validation"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		log.Printf("Error trying to marshal object, error=%s", err.Error())
		errRest := validation.ValidateUserError(err)

		c.JSON(http.StatusBadRequest, errRest)
		return
	}

	fmt.Println(userRequest)
}
