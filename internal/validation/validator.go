package validation

import (
	"errors"

	urlRepository "github.com/farpat/go-url-shortener/internal/repositories"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()

	validate.RegisterValidation("unique_slug", func(fl validator.FieldLevel) bool {
		_, err := urlRepository.Find(fl.Field().String())
		var notFoundError *urlRepository.NotFoundError

		// error is "NotFoundError" => slug does not exist => good
		return errors.As(err, &notFoundError)
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
