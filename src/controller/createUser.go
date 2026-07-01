package controller

import (
	"net/http"

	"github.com/Maryszxxx/gocrud.git/src/controller/model"
	"github.com/Maryszxxx/gocrud.git/src/controller/model/view/test/config/logger"
	"github.com/Maryszxxx/gocrud.git/src/controller/model/view/test/config/rest_err/request"
	"github.com/Maryszxxx/gocrud.git/src/controller/model/view/test/config/validation"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// UserDomainInterface é uma variável global que guarda a interface de domínio
// do usuário. Hoje ela não está sendo usada dentro de CreateUser (que usa
// model.NewUserDomain diretamente), mas normalmente serviria para permitir
// trocar a implementação real por um mock em testes (injeção de dependência).
var (
	UserDomainInterface model.UserDomainInterface
)

// CreateUser é o controller (camada HTTP) responsável por criar um usuário.
// Fluxo: recebe o JSON -> valida -> cria o "domain" -> chama a regra de negócio -> responde.
func CreateUser(c *gin.Context) {
	// Loga o início da execução dessa "jornada" (journey), útil para
	// rastrear logs relacionados à mesma operação.
	logger.Info("CreateUser called",
		zap.String("journey", "CreateUser"),
	)

	// userRequest vai receber os dados enviados no corpo (body) da requisição.
	var userRequest request.UserRequest

	// ShouldBindJSON faz o parse do JSON do body para a struct userRequest
	// e AUTOMATICAMENTE aplica as validações definidas nas tags `binding`
	// da struct UserRequest (ex: required, min, max, email...).
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		// Se o JSON for inválido ou não passar nas validações, loga o erro...
		logger.Error("Error binding JSON: ", err)

		// ...e traduz esse erro (do pacote validator) para um formato de erro
		// da API mais amigável (RestErr), usando a função do pacote validation.
		errRest := validation.ValidateUserError(err)

		// Responde com HTTP 400 (Bad Request) e o corpo do erro em JSON.
		c.JSON(http.StatusBadRequest, errRest)
		return
	}

	// Se a validação passou, cria uma entidade de domínio (UserDomain) a partir
	// dos dados recebidos. NewUserDomain retorna a interface UserDomainInterface.
	domain := model.NewUserDomain(
		userRequest.Name,
		userRequest.Password,
		userRequest.Email,
		userRequest.Age,
	)

	// Chama a regra de negócio de criação do usuário (hoje ela só faz o
	// hash da senha; a persistência em banco ainda está pendente).
	// Se CreateUser() retornar um erro, ele já vem no formato RestErr,
	// com o Code correspondente ao status HTTP.
	if err := domain.CreateUser(); err != nil {
		c.JSON(int(err.Code), err)
		return
	}

	logger.Info("User created successfully",
		zap.String("journey", "CreateUser"))

	// Responde com HTTP 200 e o próprio domain serializado em JSON.
	// Atenção: como UserDomain não tem tags `json:"-"` na senha, o hash
	// da senha também é devolvido na resposta (ver observação abaixo).
	c.JSON(http.StatusOK, domain)
}
