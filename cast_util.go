package ohanakoutilgo

import (
	"fmt"
	"github.com/ohanakogo/errhandlergo"
	"math"
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
	// Define a helper function for casting to integers
	castIntNumber := func(MIN any, MAX any, signed bool) {
		if signed {
			CastThen[int64](obj, func(i int64) {
				if CastTo[int64](MIN) < i && i < CastTo[int64](MIN) {
					result = T(obj)
				} else {
					result = 0
				}
			})
		} else {
			CastThen[uint64](obj, func(i uint64) {
				if CastTo[uint64](MIN) <= i && i <= CastTo[uint64](MIN) {
					result = T(obj)
				} else {
					result = 0
				}
			})
		}
	}

	// Check for integer types
	switch {
	case Is[int](obj):
		castIntNumber(math.MinInt, math.MaxInt, true)
	case Is[int8](obj):
		castIntNumber(math.MinInt8, math.MaxInt8, true)
	case Is[int16](obj):
		castIntNumber(math.MinInt16, math.MaxInt16, true)
	case Is[int32](obj):
		castIntNumber(math.MinInt32, math.MaxInt32, true)
	case Is[int64](obj):
		castIntNumber(math.MinInt64, math.MaxInt64, true)
	case Is[uint](obj):
		castIntNumber(0, math.MaxUint, false)
	case Is[uint8](obj):
		castIntNumber(0, math.MaxUint8, false)
	case Is[uint16](obj):
		castIntNumber(0, math.MaxUint16, false)
	case Is[uint32](obj):
		castIntNumber(0, math.MaxUint32, false)
	case Is[uint64](obj):
		castIntNumber(0, math.MaxUint64, false)
	}

	// Check for floating point types
	if Is[float32](obj) || Is[float64](obj) {
		return T(obj)
	}

	return 0
}
