package iter

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

func Times(n int, fn func(int)) {
	for i := 0; i < n; i++ {
		fn(i)
	}
}
