package collection

import (
	"fmt"
	"reflect"
)

func PluckUint64(list interface{}, fieldName string) []uint64 {
	var result []uint64
	vo := reflect.ValueOf(list)
	switch vo.Kind() {
	case reflect.Array, reflect.Slice:
		for i := 0; i < vo.Len(); i++ {
			elem := vo.Index(i)

			for elem.Kind() == reflect.Ptr {
				elem = elem.Elem()
			}

			// if nil element, skip
			if !elem.IsValid() {
				continue
			}

			if elem.Kind() != reflect.Struct {
				panic("element not struct")
			}

			f := elem.FieldByName(fieldName)
			if !f.IsValid() {
				panic(fmt.Sprintf("struct missed field %s", fieldName))
			}

			if f.Kind() != reflect.Uint64 {
				panic(fmt.Sprintf("struct element %s type required uint64", fieldName))
			}

			result = append(result, f.Uint())
		}
	default:
		panic("required list of struct type")
	}
	return result
}
