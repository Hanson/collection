package collection

import (
	"fmt"
	"reflect"
)

func PluckUint64(list interface{}, fieldName string) Uint64s {
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

func PluckString(list interface{}, fieldName string) []string {
	var result []string
	vo := reflect.ValueOf(list)
	switch vo.Kind() {
	case reflect.Array, reflect.Slice:
		for i := 0; i < vo.Len(); i++ {
			elem := vo.Index(i)

			for elem.Kind() == reflect.Ptr {
				elem = elem.Elem()
			}

			// 如果某一个元素的nil，跳过从这个元素中获取数据
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

			if f.Kind() != reflect.String {
				panic(fmt.Sprintf("struct element %s type required string", fieldName))
			}
			result = append(result, f.String())
		}
	default:
		panic("required list of struct type")
	}
	return result
}
