package ohanakoutilgo

import "reflect"

// TypeOf get the type of the generic type
func TypeOf[T any]() reflect.Type {
	return reflect.TypeOf((*T)(nil)).Elem()
}

// ActualTypeOf recursively get the non-pointer type of the generic type
func ActualTypeOf[T any]() reflect.Type {
	vType := TypeOf[T]()
	for {
		if vType.Kind() != reflect.Ptr {
			break
		}
		vType = vType.Elem()
	}
	return vType
}
