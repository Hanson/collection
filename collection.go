package collection

import (
	"fmt"
	"reflect"
)

func KeyBy(list interface{}, fieldName string) interface{} {
	lv := reflect.ValueOf(list)

	switch lv.Kind() {
	case reflect.Slice, reflect.Array:
	default:
		panic("list required slice or array type")
	}

	ev := lv.Type().Elem()
	evs := ev
	for evs.Kind() == reflect.Ptr {
		evs = evs.Elem()
	}

	if evs.Kind() != reflect.Struct {
		panic("element not struct")
	}

	field, ok := evs.FieldByName(fieldName)
	if !ok {
		panic(fmt.Sprintf("field %s not found", fieldName))
	}

	m := reflect.MakeMapWithSize(reflect.MapOf(field.Type, ev), lv.Len())
	for i := 0; i < lv.Len(); i++ {
		elem := lv.Index(i)
		elemStruct := elem
		for elemStruct.Kind() == reflect.Ptr {
			elemStruct = elemStruct.Elem()
		}

		// if nil element, skip
		if !elemStruct.IsValid() {
			continue
		}

		if elemStruct.Kind() != reflect.Struct {
			panic("element not struct")
		}

		m.SetMapIndex(elemStruct.FieldByIndex(field.Index), elem)
	}

	return m.Interface()
}

func KeyBySlice(list interface{}, fieldName string) interface{} {
	lv := reflect.ValueOf(list)

	switch lv.Kind() {
	case reflect.Slice, reflect.Array:
	default:
		panic("list required slice or array type")
	}

	lt := reflect.TypeOf(list)

	ev := lv.Type().Elem()
	evs := ev
	for evs.Kind() == reflect.Ptr {
		evs = evs.Elem()
	}

	if evs.Kind() != reflect.Struct {
		panic("element not struct")
	}

	field, ok := evs.FieldByName(fieldName)
	if !ok {
		panic(fmt.Sprintf("field %s not found", fieldName))
	}

	m := reflect.MakeMapWithSize(reflect.MapOf(field.Type, lt), lv.Len())
	for i := 0; i < lv.Len(); i++ {
		elem := lv.Index(i)
		elemStruct := elem
		for elemStruct.Kind() == reflect.Ptr {
			elemStruct = elemStruct.Elem()
		}

		// if nil element, skip
		if !elemStruct.IsValid() {
			continue
		}

		if elemStruct.Kind() != reflect.Struct {
			panic("element not struct")
		}

		vs := m.MapIndex(elemStruct.FieldByIndex(field.Index))
		if vs.IsValid() {

			vs = reflect.Append(vs, elem)
			m.SetMapIndex(elemStruct.FieldByIndex(field.Index), vs)
		} else {
			vs := reflect.MakeSlice(lt, 0, 0)
			vs = reflect.Append(vs, elem)
			m.SetMapIndex(elemStruct.FieldByIndex(field.Index), vs)
		}

	}

	return m.Interface()
}
