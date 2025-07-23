package common

import (
	"fmt"
	"reflect"
	"slices"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

// convert string into specfic supplied value
// but still return as any
func ConvertStringByType(initialValue string, datatype string) any {

	// init error variable
	var err error

	// init result
	var value any

	// check if type is allowed
	if !slices.Contains(ConvertAllowedTypes, datatype) {
		fmt.Println(fmt.Errorf("wrong types on unique validation. type: %s, data: %v", datatype, initialValue))
		return false
	}

	// switch case
	switch datatype {
	case "string":
		value = initialValue

	case "int", "int8", "int16", "int32", "int64":
		mutatedValue, err := strconv.Atoi(initialValue)

		if err != nil {
			fmt.Println(fmt.Printf("Value cannot be converted into %s. value: %s", datatype, initialValue))
			return false
		}

		switch datatype {
		case "int8":
			value = int8(mutatedValue)
		case "int16":
			value = int16(mutatedValue)
		case "int32":
			value = int32(mutatedValue)
		case "int64":
			value = int64(mutatedValue)
		default:
			value = mutatedValue
		}

	case "uint", "uint8", "uint16", "uint32", "uint64":
		mutatedValue, err := strconv.Atoi(initialValue)
		if err != nil {
			fmt.Println(fmt.Printf("Value cannot be converted into %s. value: %s", datatype, initialValue))
			return false
		}
		switch datatype {
		case "int8":
			value = uint8(mutatedValue)
		case "int16":
			value = uint16(mutatedValue)
		case "int32":
			value = uint32(mutatedValue)
		case "int64":
			value = uint64(mutatedValue)
		default:
			value = uint(mutatedValue)
		}

	case "float32", "float64":
		size, _ := strconv.Atoi(strings.ReplaceAll(datatype, "float", ""))
		value, err = strconv.ParseFloat(initialValue, size)
		if err != nil {
			fmt.Println(fmt.Printf("Value cannot be converted into %s. value: %s", datatype, initialValue))
			return false
		}

	case "bool":
		value, err = strconv.ParseBool(initialValue)
		if err != nil {
			fmt.Println(fmt.Printf("Value cannot be converted into %s. value: %s", datatype, initialValue))
			return false
		}
	}

	// return
	return value
}

// convert reflect.value that could be anything into a specific value
// but still return as any
func ConvertFieldValueByType(field reflect.Value) any {
	// field
	datatype := field.Type().String()

	// set tempValue
	var tempValue any

	if datatype[0:4] == "uint" {
		tempValue = field.Uint()
	} else if datatype[0:3] == "int" {
		tempValue = field.Int()
	} else if datatype == "bool" {
		tempValue = field.Bool()
	} else if datatype == "string" {
		tempValue = field.String()
	} else if datatype[0:5] == "float" {
		tempValue = field.Float()
	} else {
		tempValue = field.String()
	}

	// all must can be converted into string
	initialValue := fmt.Sprintf("%v", tempValue)

	// init result
	var value any

	// check if type is allowed
	if !slices.Contains(ConvertAllowedTypes, datatype) {
		fmt.Println(fmt.Errorf("wrong types on unique validation. type: %s, data: %v", datatype, initialValue))
		return false
	}

	// value
	value = ConvertStringByType(initialValue, datatype)

	// return
	return value
}

// Convert error from validation that was a string into a structured
// json format with customizable error response
func ConvertValidationError(errorText string, errorResponse map[string]ErrorMessageInterface) map[string][]string {

	// create result
	var result map[string][]string = make(map[string][]string)

	// split errors by new line
	errors := strings.SplitSeq(errorText, "\n")

	// iterate errors
	for errorItem := range errors {

		// needed phrases
		var neededPhrases []string

		// split phrases and parse only the item that contain ' or single-quote
		phrases := strings.SplitSeq(errorItem, " ")
		for phrase := range phrases {
			if strings.Contains(phrase, "'") {
				neededPhrases = append(neededPhrases, phrase)
			}
		}

		// create new variable
		key := strings.ReplaceAll(neededPhrases[0], "'", "")
		input := strings.ReplaceAll(neededPhrases[1], "'", "")
		validation := strings.ReplaceAll(neededPhrases[2], "'", "")

		// var messages and field
		var field string
		var message string

		// check if response exist
		response, responseExist := errorResponse[input]
		if !responseExist {

			// find another response
			keyResponse, keyResponseExist := errorResponse[key]

			if keyResponseExist {

				field = keyResponse.Field
				message = keyResponse.Messages[validation]

			} else {

				// put default input and message
				field = key
				message = "Form tidak valid"
			}

		} else {

			// set field
			field = response.Field

			// get message and check if exist
			var messageExist bool
			message, messageExist = response.Messages[validation]
			if !messageExist {
				message = "Form tidak valid"
			}
		}

		// check if already exist
		_, resultInputExist := result[field]

		// if exist
		if resultInputExist {

			result[field] = append(result[field], message)

		} else { // if not exist

			result[field] = []string{message}
		}
	}

	// return
	return result
}

// process index
func ProcessIndexQuery(db *gorm.DB, queries []IndexQuery, tableName string, aliases map[string]string) {
	// set variable
	var queryString string
	var queryColumn string
	var queryValue any

	// switch case
	for _, query := range queries {

		// mutate query column
		alias, aliasExist := aliases[queryColumn]
		if aliasExist {
			queryColumn = alias
		} else {
			queryColumn = fmt.Sprintf("%s.%s", tableName, query.QueryColumn)
		}

		// mutate query value
		if strings.ToLower(query.QueryValue) == "null" {
			queryValue = nil
		} else {
			queryValue = query.QueryValue
		}

		switch query.QueryCommand {
		case "is":
			queryString = fmt.Sprintf("%s = ?", queryColumn)
			db.Where(queryString, queryValue)
		case "is_not":
			queryString = fmt.Sprintf("%s != ?", queryColumn)
			db.Where(queryString, queryValue)
		case "contain":
			queryString = fmt.Sprintf("%s LIKE ?", queryColumn)
			db.Where(queryString, "%"+fmt.Sprintf("%v", queryValue)+"%")
		case "not_contain":
			queryString = fmt.Sprintf("%s NOT LIKE ?", queryColumn)
			db.Where(queryString, "%"+fmt.Sprintf("%v", queryValue)+"%")
		}
	}
}
