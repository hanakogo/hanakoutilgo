package hanakoutilgo

import (
	"fmt"
	"github.com/hanakogo/errhandlergo"
	"reflect"
)

// CastTo casts an object to a given type T. If the cast is successful, returns the value of the cast object.
// Otherwise, panics with a message indicating that the object cannot be cast as type T.
func CastTo[T any](obj any) T {
	if val, ok := obj.(T); ok {
		return val
	}
	panic(fmt.Sprintf("can't cast object as type<%s>", TypeOf[T]()))
}

// CastThen calls a function `then` with the result of casting an object `obj` to type T.
// If the cast fails, silently recovers from the panic using `errhandlergo`.
func CastThen[T any](obj any, then func(T)) {
	defer errhandlergo.HandleRecover(func(err any) {})
	then(CastTo[T](obj))
}

// Is checks if an object can be cast to a given type T.
// If the cast succeeds, returns true. Otherwise, returns false.
// Silently recovers from the panic using `errhandlergo`.
func Is[T any](obj any) (result bool) {
	defer errhandlergo.HandleRecover(func(err any) {
		result = false
	})
	CastTo[T](obj)
	return true
}

// CastToString casts an object to a string type.
// If the cast fails, returns an empty string.
// Silently recovers from the panic using `errhandlergo`.
func CastToString(obj any) (result string) {
	defer errhandlergo.HandleRecover(func(err any) {
		if val, ok := obj.(fmt.Stringer); ok {
			result = val.String()
			return
		}
		result = ""
	})
	result = CastTo[string](obj)
	return
}

// CastToNumber cast to any number type, default 0
func CastToNumber[T int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64](obj any) (result T) {
	kinds := []reflect.Kind{
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64,
	}

	rType := TypeOf[T]()
	vKind := reflect.TypeOf(obj).Kind()

	for _, kind := range kinds {
		if kind == vKind {
			convertValue := reflect.ValueOf(obj).Convert(rType)
			val := convertValue.Interface()
			result = CastTo[T](val)
			break
		}
	}

	return
}
