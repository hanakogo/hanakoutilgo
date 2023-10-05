package hanakoutilgo

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

// IsPointer check value is pointer or not
func IsPointer(v any) bool {
	if v == nil {
		return true
	}
	return reflect.TypeOf(v).Kind() == reflect.Ptr
}

// IsNil check value is nil or not
func IsNil(v any) bool {
	if !IsPointer(v) {
		return false
	}
	if v == nil {
		return true
	}
	return reflect.ValueOf(v).IsNil()
}
