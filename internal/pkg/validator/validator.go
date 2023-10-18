package validator

import (
	"errors"
	"fmt"

	"github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewCustomValidator),
	fx.Provide(NewValidator),
)

type (
	ValidationHandler struct {
		Tag          string
		Func         validator.Func
		ErrorMessage string
	}

	Validator struct {
		Validator           *validator.Validate
		ValidatorHandlerMap map[string]ValidationHandler
	}

	Params interface {
		Validate() error
	}
)

func NewCustomValidator() (configs []ValidationHandler) {
	return
}

func NewValidator(customValidators []ValidationHandler) *Validator {
	validatorPkg := validator.New()

	validatorHandlerMap := make(map[string]ValidationHandler)
	for _, customValidator := range customValidators {
		validatorHandlerMap[customValidator.Tag] = customValidator
		_ = validatorPkg.RegisterValidation(customValidator.Tag, customValidator.Func)
	}

	return &Validator{
		Validator:           validatorPkg,
		ValidatorHandlerMap: validatorHandlerMap,
	}
}

func (v *Validator) translate() (trans ut.Translator) {
	var (
		locale     = "id"
		indonesian = id.New()
		uni        = ut.New(indonesian, indonesian)
	)

	trans, _ = uni.GetTranslator(locale)
	return
}

func (v *Validator) Validate(params interface{}) (err error) {
	if err = v.Validator.Struct(params); err != nil {
		var validationErrors validator.ValidationErrors
		ok := errors.As(err, &validationErrors)

		if ok {
			for _, err := range validationErrors {
				message := err.Translate(v.translate())

				return fmt.Errorf("field: %s, error: %s", err.StructField(), message)
			}
		}
	}

	return
}
