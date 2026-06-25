package internal

import (
	"fmt"
	"reflect"
)

func ComparePtr[T comparable](fieldName string, expected, result *T) error {
	// Check if T is a struct
	if reflect.TypeOf(*new(T)).Kind() == reflect.Struct {
		if expected == nil || result == nil {
			if expected != result {
				return fmt.Errorf("%s: expected %#v but got %#v", fieldName, expected, result)
			}
		}
	} else {
		// For non-struct types, perform value comparison
		if expected == nil || result == nil {
			if expected != result {
				return fmt.Errorf("%s: expected %#v but got%#v", fieldName, expected, result)
			}
		} else if *expected != *result {
			return fmt.Errorf("%s: expected %#v but got %#v", fieldName, expected, result)
		}
	}
	return nil
}
