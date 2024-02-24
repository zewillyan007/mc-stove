package helper

import (
	"errors"
	"reflect"
)

func StructMapByFieldName(src, dest interface{}) error {

	if reflect.TypeOf(src).Kind() != reflect.Ptr || reflect.TypeOf(dest).Kind() != reflect.Ptr {
		return errors.New("src and dst must be addressable.")
	}
	dic := make(map[string]reflect.Value)
	srcPtr := reflect.ValueOf(src).Elem()
	destPtr := reflect.ValueOf(dest).Elem()
	for i := 0; i < srcPtr.NumField(); i++ {
		field := srcPtr.Type().Field(i)
		dic[field.Name] = srcPtr.FieldByName(field.Name)
	}
	for i := 0; i < destPtr.NumField(); i++ {
		currentField := destPtr.Type().Field(i)
		name := currentField.Name
		if dic[name].IsValid() && dic[name].Kind() == currentField.Type.Kind() && dic[name].CanSet() {
			destPtr.FieldByName(name).Set(dic[name])
		}
	}
	return nil
}

func StructMapByFieldTag(src, dest interface{}) error {

	if reflect.TypeOf(src).Kind() != reflect.Ptr || reflect.TypeOf(dest).Kind() != reflect.Ptr {
		return errors.New("src and dst must be addressable.")
	}
	tSrc, vSrc, tDst, vDst := reflect.TypeOf(src).Elem(), reflect.ValueOf(src).Elem(), reflect.TypeOf(dest).Elem(), reflect.ValueOf(dest).Elem()
	tagMap := make(map[string]reflect.Value)
	for i := 0; i < vDst.NumField(); i++ {
		if val, ok := tDst.Field(i).Tag.Lookup("mapper"); ok {
			tagMap[val] = vDst.Field(i)
		}
	}
	for i := 0; i < vSrc.NumField(); i++ {
		if val, ok := tSrc.Field(i).Tag.Lookup("mapper"); ok {
			if value, ok := tagMap[val]; ok && value.IsValid() && value.CanSet() && vSrc.Field(i).Kind() == value.Kind() {
				value.Set(vSrc.Field(i))
			}
		}
	}
	return nil
}
