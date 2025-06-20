package flow

import "github.com/thegogod/cmark/reflect"

type ScopeEntry struct {
	Type  reflect.Type
	Value reflect.Value
}
