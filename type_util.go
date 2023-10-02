package ohanakoutilgo

import "reflect"

func TypeOf[T any]() reflect.Type {
	return reflect.TypeOf((*T)(nil))
}

func ActualTypeOf[T any]() reflect.Type {
	vType := TypeOf[T]()
	if vType.Kind() == reflect.Ptr {
		vType = vType.Elem()
	}
	return vType
}
