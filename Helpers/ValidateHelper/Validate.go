package ValidateHelper

import (
	enLocalePackage "github.com/go-playground/locales/en"
	universalTranslatorPackage "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslationsPackage "github.com/go-playground/validator/v10/translations/en"
)

// Validator
var (
	universal *universalTranslatorPackage.UniversalTranslator
	validate  *validator.Validate
)

// Validate structure
func Validate(validateStruct interface{}) (errs validator.ValidationErrorsTranslations, err error) {
	en := enLocalePackage.New()
	universal = universalTranslatorPackage.New(en, en)
	translator, _ := universal.GetTranslator("en")

	validate = validator.New()
	if err = enTranslationsPackage.RegisterDefaultTranslations(validate, translator); err != nil {
		return
	}

	validateErr := validate.Struct(validateStruct)
	if validateErr != nil {
		errs := validateErr.(validator.ValidationErrors)
		return errs.Translate(translator), err
	}

	return
}
