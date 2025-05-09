package validation

import (
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()

	validate.RegisterValidation("unique_slug", ValidateUniqueSlug)
}

func FormatErrors(errs validator.ValidationErrors) map[string]string {
	if errs == nil {
		return nil
	}

	messages := make(map[string]string)
	for _, e := range errs {
		messages[e.Field()] = e.Tag()
	}
	return messages
}

func GetValidate() *validator.Validate {
	return validate
}
