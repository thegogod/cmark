package tx

type Tx interface {
	Rollback()
}
