package converter

import (
	"github.com/Maryszxxx/gocrud.git/src/model"
	"github.com/Maryszxxx/gocrud.git/src/model/repository/entity"
)

func ConvertDomainToEntity(
	domain model.UserDomainInterface,
) *entity.UserEntity {
	return &entity.UserEntity{
		Email:    domain.GetEmail(),
		Password: domain.GetPassword(),
		Name:     domain.GetName(),
		Age:      domain.GetAge(),
	}
}
