package maps

import "slices"

type OMap[K comparable, V any] []KeyValue[K, V]

func (self OMap[K, V]) Exists(key K) bool {
	return slices.ContainsFunc(self, func(pair KeyValue[K, V]) bool {
		return pair.Key == key
	})
}

func (self OMap[K, V]) Get(key K) V {
	i := slices.IndexFunc(self, func(pair KeyValue[K, V]) bool {
		return pair.Key == key
	})

	return self[i].Value
}

func (self OMap[K, V]) GetOrDefault(key K) V {
	var value V

	i := slices.IndexFunc(self, func(pair KeyValue[K, V]) bool {
		return pair.Key == key
	})

	if i > -1 {
		value = self[i].Value
	}

	return value
}

func (self *OMap[K, V]) Set(key K, value V) {
	arr := *self
	i := slices.IndexFunc(*self, func(pair KeyValue[K, V]) bool {
		return pair.Key == key
	})

	if i > -1 {
		arr[i].Value = value
	} else {
		*self = append(*self, KeyValue[K, V]{
			Key:   key,
			Value: value,
		})
	}
}

func (self *OMap[K, V]) Del(key K) {
	*self = slices.DeleteFunc(*self, func(pair KeyValue[K, V]) bool {
		return pair.Key == key
	})
}

func (self *OMap[K, V]) Merge(other OMap[K, V]) {
	for _, pair := range other {
		self.Set(pair.Key, pair.Value)
	}
}
