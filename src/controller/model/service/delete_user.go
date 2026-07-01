package service

import "github.com/Maryszxxx/gocrud.git/src/controller/model/view/test/config/rest_err"

// DeleteUser implementa o método DeleteUser(string) da interface UserDomainInterface.
// Recebe o "id" do usuário a ser removido.
// AINDA NÃO IMPLEMENTADO: aqui entraria a lógica de deletar o registro
// correspondente no banco de dados (ex: db.DeleteOne(...)), retornando
// um *rest_err.RestErr caso o usuário não seja encontrado ou ocorra
// algum erro de banco.
func (ud *userDomainService) DeleteUser(id string) *rest_err.RestErr {
	// lógica de deletar no banco
	return nil
}
