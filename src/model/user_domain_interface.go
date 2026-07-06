package model

import "github.com/Maryszxxx/gocrud.git/src/config/rest_err"

type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetAge() int8
	GetName() string
	GetID() string

	SetID(string)

	EncryptPassword()
	GenerateToken() (string, *rest_err.RestErr)
}
