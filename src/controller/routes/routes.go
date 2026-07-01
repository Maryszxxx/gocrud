package routes

import (
	"github.com/Maryszxxx/gocrud.git/src/controller"
	"github.com/gin-gonic/gin"
)

// InitRoutes recebe o RouterGroup do Gin (vindo do main.go) e registra
// todas as rotas da API, associando cada URL + método HTTP a uma função
// do pacote controller que vai tratar aquela requisição.
//
// É a "tabela de roteamento" da aplicação: só diz PARA ONDE cada
// requisição deve ir, sem conter nenhuma lógica de negócio.
func InitRoutes(r *gin.RouterGroup) {
	// GET /getUserById/:userId  -> busca um usuário pelo ID passado na URL
	// (":userId" é um parâmetro de rota, acessado dentro do controller via c.Param("userId"))
	r.GET("/getUserById/:userId", controller.FindUserById)

	// GET /getUserByEmail/:email -> busca um usuário pelo e-mail passado na URL
	r.GET("/getUserByEmail/:email", controller.FindUserByEmail)

	// POST /createUser -> cria um novo usuário. Espera um corpo JSON.
	r.POST("/createUser", controller.CreateUser)

	// PUT /updateUser/:userId -> atualiza os dados de um usuário existente
	r.PUT("/updateUser/:userId", controller.UpdateUser)

	// DELETE /deleteUser/:userId -> remove um usuário pelo ID
	r.DELETE("/deleteUser/:userId", controller.DeleteUser)

}
