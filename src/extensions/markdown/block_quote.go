package markdown

import (
	"github.com/thegogod/cmark/html"
	"github.com/thegogod/cmark/tokens"
)

func (self *Markdown) ParseBlockQuote(parser html.Parser, ptr *tokens.Pointer) (html.Node, error) {
	return self.parseBlockQuote(parser, NewScanner(ptr))
}

func (self *Markdown) parseBlockQuote(parser html.Parser, scan *_Scanner) (*html.BlockQuoteElement, error) {
	if !scan.Match(GreaterThan) {
		return nil, scan.Curr().Error("expected '>'")
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

		if scan.Curr().Kind() != GreaterThan {
			break
		}
	}

	self.blockQuoteDepth--
	return el, nil
}
