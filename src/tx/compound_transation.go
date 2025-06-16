package tx

type CompoundTransaction []Tx

func Compound(tx ...Tx) CompoundTransaction {
	return tx
}

func (self CompoundTransaction) Rollback() {
	for _, tx := range self {
		tx.Rollback()
	}
}
