package model

import (
	"encoding/json"
	"fmt"
)

func (ud *userDomain) SetID(id string) {
	ud.ID = id
}

type userDomain struct {
	ID       string
	email    string
	password string
	name     string
	age      int8
}

func (ud *userDomain) GetJSONValue() (string, error) {
	b, err := json.Marshal(ud)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return string(b), nil
}

func (ud *userDomain) GetID() string {
	return ud.ID
}

func (ud *userDomain) SetID(id string) {
	ud.ID = id
}

func (ud *userDomain) GetEmail() string {
	return ud.email
}

func (ud *userDomain) GetPassword() string {
	return ud.password
}

func (ud *userDomain) GetName() string {
	return ud.name
}

func (ud *userDomain) GetAge() int8 {
	return ud.age
}
