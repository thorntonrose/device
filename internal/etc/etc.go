package etc

import (
	"fmt"
	"reflect"
)

func Assert(condition bool, err error) {
	if !condition {
		panic(err)
	}
}

func Check(err error) {
	Assert(err == nil, err)
}

func Must[T any](val T, err error) T {
	Check(err)
	return val
}

func Recover(fn func(error)) {
	if r := recover(); r != nil {
		e, ok := r.(error)
		fn(If(ok, e, fmt.Errorf("%v", r)))
	}
}

//-----------------------------------------------------------------------------

func If[T any](condition bool, trueVal T, falseVal T) T {
	if condition {
		return trueVal
	}

	return falseVal
}

func Value[T any](val T, def T) T {
	return If(reflect.ValueOf(val).IsZero(), def, val)
}
