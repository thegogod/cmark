package markdown

import (
	"github.com/thegogod/cmark/ast"
	"github.com/thegogod/cmark/html"
	"github.com/thegogod/cmark/tokens"
)

func (self *Markdown) ParseUnorderedList(parser ast.Parser, ptr *tokens.Pointer) (ast.Node, error) {
	return self.parseUnorderedList(parser, NewScanner(ptr))
}

func (self *Markdown) parseUnorderedList(parser ast.Parser, scan *_Scanner) (*html.UnorderedListElement, error) {
	if !scan.Match(Dash) || !scan.Match(Space) {
		return nil, scan.curr.Error("expected '- '")
	}

	self.listDepth++
	ul := html.Ul()

	for {
		node, err := self.parseListItem(parser, scan)

		if node == nil || err != nil {
			self.listDepth--
			return ul, err
		}

		ul.Push(node)

		if !scan.MatchCount(Tab, self.listDepth-1) {
			break
		}

		if !(scan.Match(Dash) && scan.Match(Space)) {
			break
		}
	}

	self.listDepth--
	return ul, nil
}
