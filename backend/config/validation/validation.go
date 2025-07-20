package validation

import (
	"backend/library/customValidator"
	"backend/library/customValidator/passwordValidator"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func Register() {

	// get gin binding validator
	validation := binding.Validator.Engine().(*validator.Validate)

	// add custom validation
	// validation := validator.New()
	validation.RegisterValidation("alphanumeric_dash", customValidator.AlphaNumericDash)
	validation.RegisterValidation("is_unique", customValidator.IsUnique)
	validation.RegisterValidation("is_not_unique", customValidator.IsNotUnique)
	validation.RegisterValidation("confirmed", passwordValidator.Confirmed)
	validation.RegisterValidation("has_uppercase", passwordValidator.HasUppercase)
	validation.RegisterValidation("has_lowercase", passwordValidator.HasLowercase)
	validation.RegisterValidation("has_symbol", passwordValidator.HasSymbol)
	validation.RegisterValidation("has_number", passwordValidator.HasNumber)
}
