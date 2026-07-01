package response

// UserResponse define o formato de saída (o que a API devolve para o
// cliente) ao expor os dados de um usuário.
//
// OBSERVAÇÃO: essa struct hoje não está sendo usada em nenhum controller
// (ex: CreateUser devolve o próprio "domain" em vez de converter para
// UserResponse). O ideal seria os controllers converterem UserDomain
// para UserResponse antes de responder, para não vazar o campo Password
// (mesmo em hash) na resposta da API.
//
// Repare também que as tags `binding` aqui não têm efeito de validação de
// entrada, porque essa struct não é usada com ShouldBindJSON — binding
// tags só validam dados de REQUEST (entrada), não de response (saída).
type UserResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Age   int8   `json:"age" binding:"required"`
}
