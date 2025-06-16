package tx

type Transaction[T any] struct {
	save T
	ref  *T
}

func New[T any](save *T) Transaction[T] {
	return Transaction[T]{
		save: *save,
		ref:  save,
	}
}

func (self Transaction[T]) Rollback() {
	*self.ref = self.save
}
