package service

import (
	"github.com/Maryszxxx/gocrud.git/src/controller/model"
	"github.com/Maryszxxx/gocrud.git/src/controller/model/view/test/config/rest_err"
)

// FindUser implementa o método FindUser(string) da interface UserDomainInterface.
// Recebe o "id" (ou poderia ser usado também para busca por email, dependendo
// de como for implementado) e deveria retornar o UserDomain encontrado.
// AINDA NÃO IMPLEMENTADO: aqui entraria a busca no banco de dados,
// retornando (*UserDomain, nil) em caso de sucesso, ou
// (nil, rest_err.NewNotFoundError(...)) caso o usuário não exista.
func (ud *userDomainService) FindUser(id string) (*model.UserDomainInterface, *rest_err.RestErr) {
	// lógica de buscar no banco
	return nil, nil
}
