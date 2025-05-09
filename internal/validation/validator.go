package validation

import (
	"github.com/farpat/go-url-shortener/internal/repositories"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()

	validate.RegisterValidation("unique_slug", func(fl validator.FieldLevel) bool {
		exists, err := repositories.NewUrlRepository().Exists(fl.Field().String())
		if err != nil {
			panic(err)
		}
		return !exists
	})
}

func FormatErrors(errs validator.ValidationErrors) map[string]string {
	var messages = map[string]string{}
	for _, e := range errs {
		messages[e.Field()] = e.Tag()
	}
	return messages
}

func GetValidator() *validator.Validate {
	return validate
}
