package cache

import "reflect"

func hashable(i interface{}) bool {
	return hashableValue(reflect.ValueOf(i))
}

func hashableValue(v reflect.Value) bool {
	k := v.Kind()
	if k < reflect.Array || k == reflect.String || k == reflect.Ptr || k == reflect.UnsafePointer {
		return true
	} else if k == reflect.Struct {
		result := true
		for i := 0; i < v.NumField(); i++ {
			result = result && hashableValue(v.Field(i))
		}
		return result
	} else if k == reflect.Interface {
		return hashableValue(v.Elem())
	}
	return false
}
