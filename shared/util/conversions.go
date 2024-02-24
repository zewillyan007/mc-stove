package util

import (
	"fmt"
	"reflect"
	"strconv"
)

type Conversions struct{}

func NewConversions() *Conversions {
	return &Conversions{}
}

func (*Conversions) ArrInterfaceToArrString(intfSl []interface{}) []string {

	out := []string{}
	for _, i := range intfSl {
		out = append(out, i.(string))
	}
	return out
}

func InterfaceToString(data interface{}) string {
	switch s := data.(type) {
	case string:
		return s
	case []byte:
		return string(s)
	}

	value := reflect.ValueOf(data)
	switch value.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(value.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(value.Uint(), 10)
	case reflect.Float64:
		return strconv.FormatFloat(value.Float(), 'g', -1, 64)
	case reflect.Float32:
		return strconv.FormatFloat(value.Float(), 'g', -1, 32)
	case reflect.Bool:
		return strconv.FormatBool(value.Bool())
	}
	return fmt.Sprintf("%v", data)
}
