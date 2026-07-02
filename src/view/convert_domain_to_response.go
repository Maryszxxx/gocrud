package view

import (
	"github.com/Maryszxxx/gocrud.git/src/controller/model/response"
	"github.com/Maryszxxx/gocrud.git/src/model"
)

func ConvertDomainToResponse(userDomain model.UserDomainInterface) response.UserResponse {
	return response.UserResponse{
		ID:    "",
		Name:  userDomain.GetName(),
		Email: userDomain.GetEmail(),
		Age:   userDomain.GetAge(),
	}
}
