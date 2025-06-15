package ast

type ScopeEntryKind uint8

const (
	Invalid ScopeEntryKind = iota
	Type
	Var
)

type ScopeEntry struct {
	Kind  ScopeEntryKind
	Name  string
	Value any
}
