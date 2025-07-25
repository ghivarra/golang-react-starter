package customValidator

import (
	"backend/library/common"
	"fmt"
	"regexp"
	"slices"
	"strings"

	"github.com/go-playground/validator/v10"
)

func AlphaNumericDash(field validator.FieldLevel) bool {
	// get value
	value := field.Field().String()

	// regexp, generated using AI
	// I'm suck at regexp
	reg := regexp.MustCompile(`^[a-zA-Z0-9-]+$`)

	// return
	return reg.MatchString(value)
}

// params value should be splitted by '-' e.g "unique:table-column-excColumn-excValue"
func IsUnique(field validator.FieldLevel) bool {
	// set error
	var err error

	// get param that should be splitted by ':' or double dot
	param := field.Param()
	params := strings.Split(param, ":")

	if len(params) < 2 {
		fmt.Println(fmt.Printf("Failed to validate, wrong parameters. Parameter: %s", param))
		return false
	}

	// type
	withException := len(params) > 4

	// convert value
	value := common.ConvertFieldValueByType(field.Field())

	// check
	var isUnique bool

	// create and execute based on supplied value types
	if withException {

		// get exception value from binding
		exceptInitial := field.Parent().FieldByName(params[3])
		exceptValue := common.ConvertFieldValueByType(exceptInitial)

		// check
		isUnique, err = checkUniqueExcept(params[0], params[1], value, params[2], exceptValue)

		// if error
		if err != nil {
			fmt.Println(fmt.Printf("failed to check unique validation with exception. Reason: %v", err))
			return false
		}

	} else {

		isUnique, err = checkUnique(params[0], params[1], value)

		// if error
		if err != nil {
			fmt.Println(fmt.Printf("failed to check unique validation. Reason: %v", err))
			return false
		}
	}

	// return
	return isUnique
}

// params value should be splitted by '-' e.g "unique:table-column-excColumn-excValue"
func IsNotUnique(field validator.FieldLevel) bool {
	// get if unique
	isUnique := IsUnique(field)

	// return
	return !isUnique
}

// is in list
func InList(field validator.FieldLevel) bool {
	// get param that should be splitted by ':' or double dot
	param := field.Param()
	params := strings.Split(param, ":")

	// convert value its type and convert again into string
	initialValue := common.ConvertFieldValueByType(field.Field())
	value := fmt.Sprintf("%v", initialValue)

	// check if in list
	return slices.Contains(params, value)
}

// is not in list
func NotInList(field validator.FieldLevel) bool {
	// check if in list
	inList := InList(field)

	// return
	return !inList
}
