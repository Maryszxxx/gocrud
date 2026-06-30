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
	Validate = validator.New()
	transl   ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		unt := ut.New(en, en)
		transl, _ = unt.GetTranslator("en")
		en_translation.RegisterDefaultTranslations(val, transl)
	}
}

func ValidateUserError(validatorErr error) *rest_err.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationErr validator.ValidationErrors

	if errors.As(validatorErr, &jsonErr) {
		return rest_err.NewBadRequestError("Invalid JSON type for field")
	} else if errors.As(validatorErr, &jsonValidationErr) {
		errorsCauses := []rest_err.Causes{}

		for _, e := range validatorErr.(validator.ValidationErrors) {
			cause := rest_err.Causes{
				Field:   e.Field(),
				Message: e.Translate(transl),
			}
			errorsCauses = append(errorsCauses, cause)
		}
		return rest_err.NewBadRequestValidationError("Invalid JSON for field ", errorsCauses)
	} else {
		return rest_err.NewBadRequestError("Invalid JSON body")
	}

}
