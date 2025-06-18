package flow

type ScopeEntryKind uint8

const (
	Invalid ScopeEntryKind = iota
	TypeScope
	VarScope
)

type ScopeEntry struct {
	Kind  ScopeEntryKind
	Name  string
	Value any
}
