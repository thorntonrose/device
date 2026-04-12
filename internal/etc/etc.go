package etc

import (
	"fmt"
	"reflect"
)

func Check(err error) {
	if err != nil {
		panic(err)
	}
}

func Each[T any](list []T, fn func(T)) {
	for _, item := range list {
		fn(item)
	}
}

func EachWithIndex[T any](list []T, fn func(T, int)) {
	for i, item := range list {
		fn(item, i)
	}
}

func EachEntry[K comparable, V any](dict map[K]V, fn func(K, V)) {
	for key, value := range dict {
		fn(key, value)
	}
}

func If[T any](condition bool, trueVal T, falseVal T) T {
	if condition {
		return trueVal
	}

	return falseVal
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

func Value[T any](val T, defaultVal T) T {
	return If(reflect.ValueOf(val).IsZero(), defaultVal, val)
}
