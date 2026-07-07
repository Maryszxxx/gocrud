package controller

import (
	"github.com/Maryszxxx/gocrud.git/src/model/service"

	"github.com/gin-gonic/gin"
)

func NewUserControllerInterface(
	serviceInterface service.UserDomainService,
) UserControllerInterface {
	return &userControllerInterface{
		service: serviceInterface,
	}
}

type UserControllerInterface interface {
	FindUserByID(c *gin.Context)
	FindUserByEmail(c *gin.Context)

	DeleteUser(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	LoginUser(c *gin.Context)
}

type userControllerInterface struct {
	service service.UserDomainService
}

// Implement UserControllerInterface methods with minimal stubs.
func (u *userControllerInterface) FindUserByIDs(c *gin.Context) {

}

func (u *userControllerInterface) FindUserByEmails(c *gin.Context) {
	c.Status(501)
}

func (u *userControllerInterface) DeleteUsers(c *gin.Context) {
	c.Status(501)
}

func (u *userControllerInterface) CreateUsers(c *gin.Context) {
	c.Status(501)
}

func (u *userControllerInterface) UpdateUsers(c *gin.Context) {
	c.Status(501)
}

func (u *userControllerInterface) LoginUser(c *gin.Context) {
	c.Status(501)
}
