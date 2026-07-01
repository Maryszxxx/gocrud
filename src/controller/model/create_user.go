package model

import (
	"github.com/Maryszxxx/gocrud.git/src/controller/model/view/test/config/logger"
	"github.com/Maryszxxx/gocrud.git/src/controller/model/view/test/config/rest_err"
	"go.uber.org/zap"
)

// CreateUser implementa o método CreateUser() da interface UserDomainInterface
// para o struct UserDomain (por isso "ud *UserDomain" antes do nome da função:
// isso é um "method receiver", ou seja, essa função pertence ao UserDomain).
//
// É aqui que fica a regra de negócio de criação do usuário — sem nenhuma
// dependência do Gin ou de HTTP, só lógica pura de domínio.
func (ud *UserDomain) CreateUser() *rest_err.RestErr {
	// Transforma a senha em texto puro em um hash (ver EncryptPassword em user_domain.go).
	ud.EncryptPassword()

	// Loga que a senha foi transformada em hash, incluindo o e-mail e o
	// próprio hash gerado (isso é só para fins de debug/estudo — em um
	// projeto real, evitar logar até o hash da senha).
	logger.Info("Password hashed",
		zap.String("email", ud.Email),
		zap.String("hashedPassword", ud.Password),
	)

	// TODO: aqui entraria a lógica de persistir o usuário no banco de dados
	// (ex: inserir num MongoDB/Postgres). Por enquanto, a função só faz o
	// hash da senha e retorna nil, simulando sucesso sem de fato salvar nada.
	return nil
}
