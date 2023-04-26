package ohanakoutilgo

import (
	"fmt"
	"github.com/ohanakogo/errhandlergo"
	"reflect"
)

// CastTo casts an object to a given type T. If the cast is successful, returns the value of the cast object.
// Otherwise, panics with a message indicating that the object cannot be cast as type T.
func CastTo[T any](obj any) T {
	if val, ok := obj.(T); ok {
		return val
	}
	panic(fmt.Sprintf("can't cast object as type<%s>", reflect.TypeOf((*T)(nil)).String()))
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
		result = ""
	})
	result = CastTo[string](obj)
	return
}

// CastToNumber cast to any number type, default 0
func CastToNumber[T int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64](obj any) (result T) {
	// Set default returns value
	result = 0

	// Check for integer types
	switch {
	case Is[int](obj):
		CastThen[int](obj, func(i int) {
			result = T(i)
		})
		return
	case Is[int8](obj):
		CastThen[int8](obj, func(i int8) {
			result = T(i)
		})
		return
	case Is[int16](obj):
		CastThen[int16](obj, func(i int16) {
			result = T(i)
		})
		return
	case Is[int32](obj):
		CastThen[int32](obj, func(i int32) {
			result = T(i)
		})
		return
	case Is[int64](obj):
		CastThen[int64](obj, func(i int64) {
			result = T(i)
		})
		return
	case Is[uint](obj):
		CastThen[uint](obj, func(i uint) {
			result = T(i)
		})
		return
	case Is[uint8](obj):
		CastThen[uint8](obj, func(i uint8) {
			result = T(i)
		})
		return
	case Is[uint16](obj):
		CastThen[uint16](obj, func(i uint16) {
			result = T(i)
		})
		return
	case Is[uint32](obj):
		CastThen[uint32](obj, func(i uint32) {
			result = T(i)
		})
		return
	case Is[uint64](obj):
		CastThen[uint64](obj, func(i uint64) {
			result = T(i)
		})
		return
	}

	// Check for floating point types
	switch {
	case Is[float32](obj):
		CastThen[float32](obj, func(f float32) {
			result = T(f)
		})
		return
	case Is[float64](obj):
		CastThen[float64](obj, func(f float64) {
			result = T(f)
		})
		return
	}

	return
}
