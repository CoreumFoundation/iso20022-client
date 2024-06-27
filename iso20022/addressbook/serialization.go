package addressbook

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

// SerializeStruct produces a deterministic text representation of a struct
func SerializeStruct(p any) string {
	inputType := reflect.TypeOf(p)
	inputKind := inputType.Kind()
	inputValue := reflect.ValueOf(p)

	for inputKind == reflect.Pointer {
		inputType = inputType.Elem()
		inputKind = inputType.Kind()
		inputValue = inputValue.Elem()
	}

	if inputKind != reflect.Struct {
		return ""
	}

	keys := make([]string, 0, inputType.NumField())
	kvMap := make(map[string]string)

	for i := 0; i < inputType.NumField(); i++ {
		property := inputType.Field(i)
		kind := property.Type.Kind()

		isNilable := false
		for kind == reflect.Pointer {
			isNilable = true
			kind = property.Type.Elem().Kind()
		}

		switch kind {
		case reflect.String:
			fieldName := inputType.Field(i).Name
			field := inputValue.FieldByName(fieldName)

			if !field.CanInterface() {
				continue
			}

			fieldInterface := field.Interface()
			var value string

			switch fieldValue := fieldInterface.(type) {
			case string:
				value = strings.ToLower(strings.TrimSpace(fieldValue))
			case []string:
				list := fieldValue
				newList := make([]string, 0, len(list))

				for _, item := range list {
					item = strings.ToLower(strings.TrimSpace(item))
					if len(value) > 0 {
						newList = append(newList, item)
					}
				}

				value = strings.Join(newList, "\n")
			}

			if len(value) == 0 {
				continue
			}

			keys = append(keys, fieldName)
			kvMap[fieldName] = value
		case reflect.Struct:
			fieldName := inputType.Field(i).Name
			val := inputValue.FieldByName(fieldName)
			if isNilable && val.IsNil() {
				continue
			}

			value := SerializeStruct(val.Interface())

			if len(value) > 0 {
				keys = append(keys, fieldName)
				kvMap[fieldName] = "(" + value + ")"
			}
		default:
			continue
		}
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	res := make([]string, len(keys))
	for i, key := range keys {
		res[i] = fmt.Sprintf("%s=%s", key, kvMap[key])
	}

	return strings.Join(res, "/")
}
