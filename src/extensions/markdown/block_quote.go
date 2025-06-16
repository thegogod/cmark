package markdown

import (
	"github.com/thegogod/cmark/ast"
	"github.com/thegogod/cmark/html"
	"github.com/thegogod/cmark/tokens"
)

func (self *Markdown) ParseBlockQuote(parser ast.Parser, ptr *tokens.Pointer) (ast.Node, error) {
	return self.parseBlockQuote(parser, NewScanner(ptr))
}

func (self *Markdown) parseBlockQuote(parser ast.Parser, scan *_Scanner) (*html.BlockQuoteElement, error) {
	if !scan.Match(GreaterThan) {
		return nil, scan.curr.Error("expected '>'")
	}

	self.blockQuoteDepth++
	el := html.BlockQuote()

	for {
		node, err := parser.ParseBlock(scan.ptr)

		if node == nil || err != nil {
			self.blockQuoteDepth--
			return el, err
		}

		el.Push(node)

		if scan.curr.Kind() != GreaterThan {
			break
		}
	}

	self.blockQuoteDepth--
	return el, nil
}
