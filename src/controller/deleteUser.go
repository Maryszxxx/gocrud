package controller

import "github.com/gin-gonic/gin"

// DeleteUser é o controller responsável por remover um usuário.
// AINDA NÃO IMPLEMENTADO: o esperado aqui seria:
//  1. pegar o "userId" da URL com c.Param("userId")
//  2. chamar algo como model.NewUserDomain(...).DeleteUser(userId)
//  3. tratar o erro (se houver) e responder com c.JSON(...)
//  4. em caso de sucesso, normalmente se responde com http.StatusOK
//     ou http.StatusNoContent
func DeleteUser(c *gin.Context) {

}
