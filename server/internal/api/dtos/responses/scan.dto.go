package dto_res

import (
	"database/sql"
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

// =====================

func TransformNullStrings(input interface{}) interface{} {
	return transformValue(reflect.ValueOf(input)).Interface()
}

func isTimeType(t reflect.Type) bool {
	return t.String() == "time.Time"
}

func transformValue(val reflect.Value) reflect.Value {
	// Handle pointers
	if val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return val
		}
		val = val.Elem()
	}

	fmt.Println("----->", val.Kind())
	switch val.Kind() {

	case reflect.Struct:
		// Create a new struct of the same type
		// fmt.Println("-----> new struct : " + val.Type().Name())
		if isTimeType(val.Type()) {
			return val
		}
		newStruct := reflect.New(val.Type()).Elem()

		for i := 0; i < val.NumField(); i++ {

			fmt.Println("-->", val.Type().Field(i).Name)
			field := val.Field(i)

			// Special handling for sql.NullString
			if nullStr, ok := field.Interface().(sql.NullString); ok {
				fmt.Println("-----> sql.NullString", nullStr.String, nullStr.Valid)
				if nullStr.Valid {
					// If Valid is true, set the field to the string value
					// newString := reflect.New(field.Type()).Elem()
					newStruct.Field(i).SetString(nullStr.String)

					// newStr := reflect.ValueOf(nullStr.String)
					// newStruct.Field(i).Set(newStr)
				}
				// If not Valid, the field is simply not set in the new struct
				continue
			}

			// Recursive transformation for nested structs or pointer to structs
			if field.Kind() == reflect.Struct ||
				(field.Kind() == reflect.Ptr && field.Type().Elem().Kind() == reflect.Struct) {
				transformedField := transformValue(field)
				if !transformedField.IsZero() {
					newStruct.Field(i).Set(transformedField)
				}
				continue
			}

			// Handling slices
			if field.Kind() == reflect.Slice {
				newSlice := reflect.MakeSlice(field.Type(), 0, field.Len())
				for j := 0; j < field.Len(); j++ {
					transformedItem := transformValue(field.Index(j))
					if !transformedItem.IsZero() {
						newSlice = reflect.Append(newSlice, transformedItem)
					}
				}
				if newSlice.Len() > 0 {
					newStruct.Field(i).Set(newSlice)
				}
				continue
			}

			// For other types, copy as-is if not zero
			if !field.IsZero() {
				newStruct.Field(i).Set(field)
			}
		}

		return newStruct

	case reflect.Slice:
		// Handle slice of structs
		newSlice := reflect.MakeSlice(val.Type(), 0, val.Len())
		for i := 0; i < val.Len(); i++ {
			transformedItem := transformValue(val.Index(i))
			if !transformedItem.IsZero() {
				newSlice = reflect.Append(newSlice, transformedItem)
			}
		}
		return newSlice

	default:
		return val
	}
}
