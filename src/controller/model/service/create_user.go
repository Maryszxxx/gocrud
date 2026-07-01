package service

import (
	"fmt"

	"github.com/Maryszxxx/gocrud.git/src/controller/model"
	"github.com/Maryszxxx/gocrud.git/src/controller/model/view/test/config/rest_err"
)

func (ud *userDomainService) CreateUser(
	userDomain model.UserDomainInterface,
) *rest_err.RestErr {

	userDomain.EncryptPassword()
	fmt.Println(userDomain.GetPassword())

	return nil
}
