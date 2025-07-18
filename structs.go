package gollama

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

func (o ChatOuput) DecodeContent(v interface{}) error {
	if o.Content == "" {
		return fmt.Errorf("content is empty")
	}

	blocks := strings.Split(o.Content, "```")

	var lastJSON string
	for _, block := range blocks {
		block = strings.TrimPrefix(block, "json")
		block = strings.TrimSpace(block)
		if strings.HasPrefix(block, "{") && strings.HasSuffix(block, "}") {
			lastJSON = block
		}
	}

	if lastJSON == "" {
		return fmt.Errorf("no JSON found in content")
	}

	err := json.Unmarshal([]byte(lastJSON), v)
	if err != nil {
		return fmt.Errorf("error decoding JSON: %w", err)
	}

	return nil
}

func StructToStructuredFormat(s interface{}) StructuredFormat {
	structValue := reflect.ValueOf(s)
	structType := structValue.Type()

	properties := make(map[string]*FormatProperty)
	required := make([]string, 0)

	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)

		strType, _, err := fieldTypeToGollamaType(field.Type)
		if err != nil {
			return StructuredFormat{}
		}

		if field.Tag.Get("ignored") == "true" {
			continue
		}

		fieldName := field.Name
		if field.Tag.Get("json") != "" {
			fieldName = field.Tag.Get("json")
		}

		property := &FormatProperty{
			Type:        strType,
			Description: field.Tag.Get("description"),
		}

		switch field.Type.Kind() {
		case reflect.Struct:
			// Recursively handle nested structs
			nested := StructToStructuredFormat(reflect.New(field.Type).Elem().Interface())
			property.Type = "object"
			property.Properties = make(map[string]*FormatProperty)
			for k, v := range nested.Properties {
				property.Properties[k] = v
			}
		case reflect.Slice:
			elemType := field.Type.Elem()
			elemKind := elemType.Kind()
			itemProperty := &FormatProperty{}
			switch elemKind {
			case reflect.Struct:
				nested := StructToStructuredFormat(reflect.New(elemType).Elem().Interface())
				itemProperty.Type = "object"
				itemProperty.Properties = make(map[string]*FormatProperty)
				for k, v := range nested.Properties {
					itemProperty.Properties[k] = v
				}
			default:
				elemTypeStr, _, _ := fieldTypeToGollamaType(elemType)
				itemProperty.Type = elemTypeStr
			}
			property.Type = "array"
			property.Items = itemProperty
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
	}
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

func AnyToStructuredFormat(val interface{}) StructuredFormat {
	m, ok := val.(map[string]interface{})
	if !ok {
		return StructuredFormat{}
	}

	sf := StructuredFormat{}

	if t, ok := m["type"].(string); ok {
		sf.Type = t
	}
	if props, ok := m["properties"].(map[string]interface{}); ok {
		sf.Properties = make(map[string]*FormatProperty)
		for k, v := range props {
			if propMap, ok := v.(map[string]interface{}); ok {
				fp := &FormatProperty{}
				if typ, ok := propMap["type"].(string); ok {
					fp.Type = typ
				}
				if desc, ok := propMap["description"].(string); ok {
					fp.Description = desc
				}
				if items, ok := propMap["items"].(map[string]interface{}); ok {
					itemProp := &FormatProperty{}
					if itemType, ok := items["type"].(string); ok {
						itemProp.Type = itemType
					}
					fp.Items = itemProp
				}
				if nestedProps, ok := propMap["properties"].(map[string]interface{}); ok {
					fp.Properties = make(map[string]*FormatProperty)
					for nk, nv := range nestedProps {
						if npMap, ok := nv.(map[string]interface{}); ok {
							nfp := &FormatProperty{}
							if ntyp, ok := npMap["type"].(string); ok {
								nfp.Type = ntyp
							}
							if ndesc, ok := npMap["description"].(string); ok {
								nfp.Description = ndesc
							}
							fp.Properties[nk] = nfp
						}
					}
				}
				sf.Properties[k] = fp
			}
		}
	}
	if req, ok := m["required"].([]interface{}); ok {
		sf.Required = make([]string, len(req))
		for i, v := range req {
			if s, ok := v.(string); ok {
				sf.Required[i] = s
			}
		}
	}

	return sf
}
