package markdown

import (
	"github.com/thegogod/cmark/html"
	"github.com/thegogod/cmark/tokens"
)

func (self *Markdown) ParseItalic(parser html.Parser, ptr *tokens.Pointer) (html.Node, error) {
	return self.parseItalic(parser, NewScanner(ptr))
}

func (self *Markdown) parseItalic(parser html.Parser, scan *Scanner) (*html.ItalicElement, error) {
	italic := html.I()

	if !scan.MatchCount(Asterisk, 1) {
		return italic, scan.Curr().Error("expected '*'")
	}

	for !scan.Match(Asterisk) {
		node, err := parser.ParseInline(scan.ptr)

		if node == nil {
			return italic, scan.Curr().Error("expected closing '*'")
		}

		if err != nil {
			return italic, err
		}

		italic.Push(node)
	}

	log.Debugln("italic")
	return italic, nil
}
