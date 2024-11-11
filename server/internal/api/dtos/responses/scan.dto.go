package dto_res

import (
	"fmt"
	"reflect"
)

func FilterStruct(v interface{}, excludeFields ...string) interface{} {
	value := reflect.ValueOf(v)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	if value.Kind() != reflect.Struct {
		return v
	}

	filteredStruct := reflect.New(value.Type()).Elem()

	for i := 0; i < value.NumField(); i++ {
		fieldName := value.Type().Field(i).Name
		fieldValue := value.Field(i)

		if containsString(excludeFields, fieldName) {
			continue
		}

		if fieldValue.Kind() == reflect.Struct {
			filteredField := FilterStruct(fieldValue.Interface(), excludeFields...)
			filteredStruct.Field(i).Set(reflect.ValueOf(filteredField))
		} else {
			filteredStruct.Field(i).Set(fieldValue)
		}
	}

	return filteredStruct.Interface()
}

func containsString(slice []string, str string) bool {
	for _, item := range slice {
		if item == str {
			return true
		}
	}
	return false
}

// =====================

func FilterStruct2(v interface{}, excludeFields ...string) interface{} {
	value := reflect.ValueOf(v)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	// Create a new struct with the same type as the input
	newStruct := reflect.New(value.Type()).Elem()

	// Iterate through the fields of the input struct
	for i := 0; i < value.NumField(); i++ {
		field := value.Type().Field(i)
		fmt.Println("----->", field.Name, field.Type.Kind() == reflect.Struct)
		fieldValue := value.Field(i)

		// Check if the field should be excluded
		shouldExclude := false
		for _, f := range excludeFields {
			if field.Name == f {
				shouldExclude = true
				break
			}
		}

		if !shouldExclude {
			// Set the field in the new struct
			newStruct.Field(i).Set(fieldValue)
		}
	}

	// If the original input was a pointer, return a pointer to the new struct
	if reflect.TypeOf(v).Kind() == reflect.Ptr {
		return newStruct.Addr().Interface()
	}

	return newStruct.Interface()
}
