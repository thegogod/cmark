package maps

type KeyValue[K comparable, V any] struct {
	Key   K
	Value V
}

func Pair[K comparable, V any](key K, value V) KeyValue[K, V] {
	return KeyValue[K, V]{
		Key:   key,
		Value: value,
	}
}
