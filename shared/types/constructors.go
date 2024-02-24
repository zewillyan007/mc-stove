package types

type FnDefaultConstructor = func(args ...interface{}) interface{}

func SetConstructor(typ interface{}, f FnDefaultConstructor) {
	defaultConstructors[typ] = f
}

func GetConstructor(typ interface{}) FnDefaultConstructor {
	return defaultConstructors[typ].(FnDefaultConstructor)
}

var defaultConstructors map[interface{}]interface{}

func init() {
	defaultConstructors = make(map[interface{}]interface{})
}
