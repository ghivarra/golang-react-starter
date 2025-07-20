package passwordValidator

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

// check if the same as password_confirmed
func Confirmed(field validator.FieldLevel) bool {
	value := field.Field().String()
	confirmation := field.Parent().FieldByName("PasswordConfirmation").String()
	return value == confirmation
}

// check if password has uppercase
func HasUppercase(field validator.FieldLevel) bool {
	value := field.Field().String()
	matched, _ := regexp.MatchString(`[A-Z]`, value)
	return matched
}

// check if password has lowercase
func HasLowercase(field validator.FieldLevel) bool {
	value := field.Field().String()
	matched, _ := regexp.MatchString(`[a-z]`, value)
	return matched
}

// check if password has symbol
func HasSymbol(field validator.FieldLevel) bool {
	value := field.Field().String()
	matched, _ := regexp.MatchString(`[^a-zA-Z0-9]`, value)
	return matched
}

// check if password has number
func HasNumber(field validator.FieldLevel) bool {
	value := field.Field().String()
	matched, _ := regexp.MatchString(`\d`, value)
	return matched
}
