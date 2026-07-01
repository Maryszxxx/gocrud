package service

import (
	"github.com/Maryszxxx/gocrud.git/src/controller/model"
	"github.com/Maryszxxx/gocrud.git/src/controller/model/view/test/config/rest_err"
)

// UpdateUser implementa o método UpdateUser(string) da interface UserDomainInterface.
// Recebe o "id" do usuário a ser atualizado; os novos dados viriam dos
// campos já preenchidos em "ud" (o próprio UserDomain que chamou o método).
// AINDA NÃO IMPLEMENTADO: aqui entraria a lógica de atualizar o registro
// no banco de dados usando o id como filtro.
func (ud *userDomainService) UpdateUser(userId string, userDomain model.UserDomainInterface) *rest_err.RestErr {
	// lógica de atualizar no banco, usando o id
	return nil
}
