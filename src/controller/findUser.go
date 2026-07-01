package controller

import "github.com/gin-gonic/gin"

// FindUserById é o controller responsável por buscar um usuário pelo ID.
// AINDA NÃO IMPLEMENTADO: o esperado aqui seria:
//  1. userId := c.Param("userId")
//  2. user, err := model.NewUserDomain("", "", "", 0).FindUser(userId)
//     (ou alguma forma de chamar a busca sem precisar instanciar um domain completo)
//  3. se err != nil, responder c.JSON(int(err.Code), err)
//  4. se sucesso, responder c.JSON(http.StatusOK, user)
func FindUserById(c *gin.Context) {

}

// FindUserByEmail é o controller responsável por buscar um usuário pelo e-mail.
// Segue a mesma lógica esperada de FindUserById, mas usando c.Param("email").
func FindUserByEmail(c *gin.Context) {

}
