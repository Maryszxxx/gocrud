package controller

import (
	"fmt"
	"net/http"

	"github.com/Maryszxxx/gocrud.git/src/controller/model/view/test/config/rest_err"
	"github.com/Maryszxxx/gocrud.git/src/controller/model/view/test/config/rest_err/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := rest_err.NewBadRequestError("invalid json body")
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, restErr)
		return
	}

	fmt.Println(userRequest)
}
