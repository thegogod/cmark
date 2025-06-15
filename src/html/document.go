package html

import (
	"strings"

	"github.com/thegogod/cmark/ast"
	"github.com/thegogod/cmark/reflect"
)

// https://developer.mozilla.org/en-US/docs/Web/API/Document
type Document []Node

// https://developer.mozilla.org/en-US/docs/Web/API/Document
func New() Document {
	return Document{}
}

func (self Document) Head() *HeadElement {
	for _, node := range self {
		if head, ok := node.(*HeadElement); ok {
			return head
		}
	}

	return nil
}

func (self Document) Body() *BodyElement {
	for _, node := range self {
		if body, ok := node.(*BodyElement); ok {
			return body
		}
	}

	return nil
}

func (self Document) Count() int {
	return len(self)
}

func (self Document) Children() []Node {
	return self
}

func (self *Document) Push(children ...Node) *Document {
	for _, child := range children {
		*self = append(*self, child)
	}

	return self
}

func (self *Document) Pop() *Document {
	arr := *self

	if len(arr) == 0 {
		return self
	}

	arr = arr[:len(arr)-1]
	*self = arr
	return self
}

func (self Document) Render(scope *ast.Scope) []byte {
	content := ""

	for _, node := range self {
		if strings.HasPrefix(node.GetSelector(), ":") {
			continue
		}

		content += string(node.Render(scope))
	}

	return []byte(content)
}

func (self Document) RenderPretty(scope *ast.Scope, indent string) []byte {
	content := []string{}

	for _, node := range self {
		if strings.HasPrefix(node.GetSelector(), ":") {
			continue
		}

		content = append(content, string(node.RenderPretty(scope, indent)))
	}

	return []byte(strings.Join(content, "\n"))
}

func (self Document) GetById(id string) Node {
	for _, child := range self {
		if node := child.GetById(id); node != nil {
			return node
		}
	}

	return nil
}

func (self Document) Select(query ...any) []Node {
	nodes := []Node{}

	for _, child := range self {
		nodes = append(nodes, child.Select(query...)...)
	}

	return nodes
}

func (self Document) Validate(scope *ast.Scope) error {
	return nil
}

func (self Document) Evaluate(scope *ast.Scope) (reflect.Value, error) {
	value := self.Render(scope)
	return reflect.NewString(string(value)), nil
}
