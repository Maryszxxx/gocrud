package controller

import "github.com/gin-gonic/gin"

// UpdateUser é o controller responsável por atualizar os dados de um usuário.
// AINDA NÃO IMPLEMENTADO: o esperado aqui seria:
//  1. userId := c.Param("userId")
//  2. fazer o bind do JSON do body para uma struct de request (igual em CreateUser)
//  3. validar os dados recebidos
//  4. montar/chamar o domain e executar domain.UpdateUser(userId)
//  5. tratar erro ou responder com sucesso (c.JSON(http.StatusOK, ...))
func UpdateUser(c *gin.Context) {

}
