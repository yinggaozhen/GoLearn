package decoder

import (
	"reflect"
	"strings"
)

func Decode(input interface{}, output interface{}) error {
	outputVal := reflect.ValueOf(output)
	inputVal := reflect.ValueOf(input)

	switch outputVal.Elem().Kind() {
	case reflect.Struct:
		decodeStruct(inputVal, outputVal)
	}

	return nil
}

func decodeStruct(inputVal reflect.Value, outputVal reflect.Value) {
	outputElem := outputVal.Elem()
	outputType := outputElem.Type()

	for i := 0; i < outputType.NumField(); i++ {
		field := outputType.Field(i)
		value := outputElem.Field(i)

		fieldName := field.Name

		if field.Tag != "" {
			if tag := field.Tag.Get("guozhen"); tag != "" {
				fieldName = tag
			}
		}

		inputFieldVal := inputVal.MapIndex(reflect.ValueOf(strings.ToLower(fieldName)))
		if inputFieldVal.IsValid() {
			if value.CanSet() {
				value.Set(reflect.ValueOf(inputFieldVal.Interface()))
			}
		}
	}
}