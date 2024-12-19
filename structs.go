package gollama

import (
	"fmt"
	"reflect"
)

func StructToStructuredFormat(s interface{}) (StructuredFormat, error) {
	structValue := reflect.ValueOf(s)
	structType := structValue.Type()

	properties := make(map[string]FormatProperty)
	required := make([]string, 0)

	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)

		strType, strItems, err := fieldTypeToGollamaType(field.Type)
		if err != nil {
			return StructuredFormat{}, err
		}

		if field.Tag.Get("ignored") == "true" {
			continue
		}

		fieldName := field.Name
		if field.Tag.Get("json") != "" {
			fieldName = field.Tag.Get("json")
		}

		property := FormatProperty{
			Type:        strType,
			Description: field.Tag.Get("description"),
		}

		if strItems != "" {
			property.Items = ItemProperty{
				Type: strItems,
			}
		}

		if field.Tag.Get("required") == "true" {
			required = append(required, fieldName)
		}

		properties[fieldName] = property
	}

	return StructuredFormat{
		Type:       "object",
		Properties: properties,
		Required:   required,
	}, nil

}

func fieldTypeToGollamaType(fieldType reflect.Type) (string, string, error) {
	switch fieldType.Kind() {
	case reflect.String:
		return "string", "", nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return "integer", "", nil
	case reflect.Float32, reflect.Float64:
		return "number", "", nil
	case reflect.Bool:
		return "boolean", "", nil
	default:
		if fieldType == reflect.TypeOf([]int{}) || fieldType == reflect.TypeOf([]int8{}) ||
			fieldType == reflect.TypeOf([]int16{}) || fieldType == reflect.TypeOf([]int32{}) ||
			fieldType == reflect.TypeOf([]int64{}) {
			return "array", "integer", nil
		} else if fieldType == reflect.TypeOf([]float32{}) || fieldType == reflect.TypeOf([]float64{}) {
			return "array", "number", nil
		} else if fieldType == reflect.TypeOf([]string{}) {
			return "array", "string", nil
		} else if fieldType == reflect.TypeOf([]bool{}) {
			return "array", "boolean", nil
		}
		return "string", "", fmt.Errorf("unsupported field type: %s", fieldType.String())
	}
}
