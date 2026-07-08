package rest_err

import "net/http"

// RestErr é o formato padronizado de erro retornado pela API.
// Usar um formato único para todos os erros facilita o consumo da API
// pelo front-end/cliente, que sempre sabe onde encontrar mensagem, código
// e causas detalhadas do erro.
type RestErr struct {
	Message string   `json:"message"` // mensagem geral do erro, legível para humanos
	Err     string   `json:"error"`   // "tipo" do erro, em formato de slug (ex: "bad_request")
	Code    int64    `json:"code"`    // código HTTP correspondente (400, 404, 500 etc)
	Causes  []Causes `json:"causes"`  // lista de causas específicas (útil em erros de validação)
}

// Causes representa uma causa específica de um erro, geralmente ligada
// a um campo em particular (ex: campo "email" -> "campo obrigatório").
type Causes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// NewRestErr é o construtor genérico: cria um RestErr com todos os
// campos definidos manualmente. As funções abaixo são "atalhos" para
// os tipos de erro HTTP mais comuns, já preenchendo Err e Code.
func NewRestErr(message string, err string, code int64, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     err,
		Code:    code,
		Causes:  causes,
	}
}

// NewBadRequestError cria um erro 400 (Bad Request) simples, sem causas detalhadas.
func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
		Causes:  []Causes{},
	}
}

func NewUnauthorizedRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "unauthorized",
		Code:    http.StatusUnauthorized,
	}
}

// NewBadRequestValidationError cria um erro 400 (Bad Request) já incluindo
// a lista de causas — usado quando a validação de campos falha (ver validateUser.go).
func NewBadRequestValidationError(message string, Causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
		Causes:  Causes,
	}
}

// NewInternalServerError cria um erro 500 (Internal Server Error),
// usado para falhas inesperadas (ex: erro de conexão com o banco).
func NewInternalServerError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "internal_server_error",
		Code:    http.StatusInternalServerError,
		Causes:  []Causes{},
	}
}

// NewNotFoundError cria um erro 404 (Not Found), usado quando o
// recurso buscado (ex: usuário por ID) não existe.
func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "not_found",
		Code:    http.StatusNotFound,
		Causes:  []Causes{},
	}
}

// NewUnauthorizedError cria um erro 401 (Unauthorized),
// usado quando falta autenticação válida.
func NewUnauthorizedError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "unauthorized",
		Code:    http.StatusUnauthorized,
		Causes:  []Causes{},
	}
}

// NewForbiddenError cria um erro 403 (Forbidden),
// usado quando o usuário está autenticado mas não tem permissão para a ação.
func NewForbiddenError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "forbidden",
		Code:    http.StatusForbidden,
		Causes:  []Causes{},
	}
}

// Error() faz o RestErr implementar a interface "error" nativa do Go
// (que exige só um método: Error() string). Isso permite, por exemplo,
// usar *RestErr em qualquer lugar que espera um "error" comum.
func (r *RestErr) Error() string {
	return r.Message
}
