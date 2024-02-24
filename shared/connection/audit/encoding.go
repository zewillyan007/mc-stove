package audit

import (
	"fmt"
	"reflect"
)

func ToData(i interface{}, tag string) (Data, error) {
	out := make(map[string]interface{})

	v := reflect.ValueOf(i)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("ToMap only accepts structs; got %T", v)
	}

	typ := v.Type()
	for i := 0; i < v.NumField(); i++ {
		fi := typ.Field(i)
		if tagName := fi.Tag.Get(tag); tagName != "" {
			out[tagName] = v.Field(i).Interface()
		}
	}
	return out, nil
}
