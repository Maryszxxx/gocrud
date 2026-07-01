package validation

import (
	"encoding/json"
	"errors"

	"github.com/Maryszxxx/gocrud.git/src/controller/model/view/test/config/rest_err"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
)

var (
	// Validate é uma instância do validador (usada aqui só como referência global,
	// já que o Gin usa internamente sua própria instância via binding.Validator).
	Validate = validator.New()

	// transl é o "tradutor" responsável por transformar mensagens de erro
	// técnicas do validator (ex: "Name is required") em mensagens já formatadas
	// no idioma configurado (inglês, neste caso).
	transl ut.Translator
)

// init() é executado automaticamente uma única vez quando o pacote é carregado
// (antes até da função main rodar). Aqui ele configura o tradutor de mensagens
// de validação em inglês.
func init() {
	// Pega o validador interno que o Gin usa para o ShouldBindJSON.
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// Cria o "locale" (idioma) inglês.
		en := en.New()
		// Cria o universal translator com inglês como idioma padrão e de fallback.
		unt := ut.New(en, en)
		// Obtém o tradutor específico para "en".
		transl, _ = unt.GetTranslator("en")
		// Registra as traduções padrão (mensagens prontas) do validator para inglês,
		// vinculando-as à instância "val" (a mesma usada pelo Gin no bind).
		en_translation.RegisterDefaultTranslations(val, transl)
	}
}

// ValidateUserError recebe o erro retornado por c.ShouldBindJSON() e o
// transforma em um erro padronizado da API (*rest_err.RestErr), tratando
// separadamente três cenários possíveis.
func ValidateUserError(validatorErr error) *rest_err.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationErr validator.ValidationErrors

	// Caso 1: o erro é de TIPO incompatível no JSON
	// (ex: mandou "age": "abc" quando o campo espera um número).
	if errors.As(validatorErr, &jsonErr) {
		return rest_err.NewBadRequestError("Invalid JSON type for field")

		// Caso 2: o erro é de VALIDAÇÃO (as tags `binding` da struct, tipo
		// required, min, max, email etc, não foram satisfeitas).
	} else if errors.As(validatorErr, &jsonValidationErr) {
		errorsCauses := []rest_err.Causes{}

		// Percorre cada erro de validação individual (pode haver mais de um
		// campo inválido ao mesmo tempo) e monta uma "causa" para cada um,
		// já com a mensagem traduzida para inglês.
		for _, e := range validatorErr.(validator.ValidationErrors) {
			cause := rest_err.Causes{
				Field:   e.Field(),
				Message: e.Translate(transl),
			}
			errorsCauses = append(errorsCauses, cause)
		}
		// Retorna um erro 400 com a lista completa de causas (um item por campo inválido).
		return rest_err.NewBadRequestValidationError("Invalid JSON for field ", errorsCauses)

		// Caso 3: qualquer outro erro de parse do JSON
		// (ex: JSON mal formado, faltando chave, vírgula sobrando etc).
	} else {
		return rest_err.NewBadRequestError("Invalid JSON body")
	}

}
