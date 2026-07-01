package request

// UserRequest define o formato de entrada esperado no corpo (body) das
// requisições que criam/alteram um usuário (ex: POST /createUser).
//
// As tags `binding` são lidas pelo Gin (através do pacote validator) quando
// c.ShouldBindJSON(&userRequest) é chamado — é isso que garante que os
// dados chegando estão dentro das regras antes de qualquer lógica rodar.
type UserRequest struct {
	// Nome: obrigatório, entre 4 e 90 caracteres.
	Name string `json:"name" binding:"required,min=4,max=90"`

	// Senha: obrigatória, entre 6 e 100 caracteres, e precisa conter:
	// - pelo menos um número (containsany=0123456789)
	// - pelo menos uma letra maiúscula (containsany=ABC...)
	// - pelo menos uma letra minúscula (containsany=abc...)
	// (ou seja, exige uma senha "forte" com número + maiúscula + minúscula)
	Password string `json:"password" binding:"required,min=6,max=100,containsany=0123456789,containsany=ABCDEFGHIJKLMNOPQRSTUVWXYZ,containsany=abcdefghijklmnopqrstuvwxyz"`

	// Email: obrigatório e precisa ter formato de e-mail válido.
	Email string `json:"email" binding:"required,email"`

	// Idade: obrigatória, entre 1 e 140.
	Age int8 `json:"age" binding:"required,min=1,max=140"`
}
