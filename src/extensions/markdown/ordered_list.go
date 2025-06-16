package markdown

import (
	"github.com/thegogod/cmark/ast"
	"github.com/thegogod/cmark/html"
	"github.com/thegogod/cmark/tokens"
)

func (self *Markdown) ParseOrderedList(parser ast.Parser, ptr *tokens.Pointer) (ast.Node, error) {
	return self.parseOrderedList(parser, NewScanner(ptr))
}

func (self *Markdown) parseOrderedList(parser ast.Parser, scan *_Scanner) (*html.OrderedListElement, error) {
	if !scan.Match(Integer) || !scan.Match(Period) || scan.Match(Space) {
		return nil, scan.Curr().Error("expected '{int}. '")
	}

	self.listDepth++
	ol := html.Ol()

	for {
		node, err := self.parseListItem(parser, scan)

		if node == nil || err != nil {
			self.listDepth--
			return ol, err
		}

		ol.Push(node)

		if !scan.MatchCount(Tab, self.listDepth-1) {
			break
		}

		if !(scan.Match(Integer) && scan.Match(Period) && scan.Match(Space)) {
			break
		}
	}

	self.listDepth--
	return ol, nil
}
