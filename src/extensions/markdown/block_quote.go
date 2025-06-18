package markdown

import (
	"github.com/thegogod/cmark/html"
	"github.com/thegogod/cmark/tokens"
)

func (self *Markdown) ParseBlockQuote(parser html.Parser, ptr *tokens.Pointer) (html.Node, error) {
	return self.parseBlockQuote(parser, NewScanner(ptr))
}

func (self *Markdown) parseBlockQuote(parser html.Parser, scan *Scanner) (*html.BlockQuoteElement, error) {
	el := html.BlockQuote()

	if !scan.Match(GreaterThan) {
		return el, scan.Curr().Error("expected '>'")
	}

	self.blockQuoteDepth++

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

	log.Debugln("block_quote")
	self.blockQuoteDepth--
	return el, nil
}
