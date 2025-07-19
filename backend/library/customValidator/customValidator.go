package customValidator

import (
	"fmt"
	"regexp"
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

// params value should be splitted by '|' e.g unique:table|column|type|excColumn|excType|excValue
func Unique(field validator.FieldLevel) bool {
	// set error
	var err error

	// get param that should be splitted by '|'
	param := field.Param()
	params := strings.Split(param, "|")

	if len(params) < 3 {
		fmt.Println(fmt.Printf("Failed to validate, wrong parameters. Parameter: %s", param))
		return false
	}

	// type
	withException := len(params) > 4

	// get value
	fieldValue := field.Field()

	// check
	var isUnique bool

	// create and execute based on supplied value types
	if withException {

		// type
		valueType := params[2]
		exceptionValueType := params[4]

		// value
		value := convertValueByType(fieldValue.String(), valueType)
		exceptionValue := convertValueByType(params[5], exceptionValueType)

		// check
		isUnique, err = checkUniqueExcept(params[0], params[1], value, params[3], exceptionValue)

		// if error
		if err != nil {
			fmt.Println(fmt.Printf("failed to check unique validation. Reason: %v", err))
			return false
		}

	} else {

		// type
		valueType := params[2]

		// value
		value := convertValueByType(fieldValue.String(), valueType)

		// check
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

// unique:string|int|int8|int16|int32|int64|uint|bool|float
