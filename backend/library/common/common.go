package common

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

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
