package validation

import (
	"github.com/farpat/go-url-shortener/internal/repositories"
	"github.com/go-playground/validator/v10"
)

func ValidateUniqueSlug(fl validator.FieldLevel) bool {
	exists, err := repositories.NewUrlRepository().Exists(fl.Field().String())
	if err != nil {
		panic(err)
	}
	return !exists
}
