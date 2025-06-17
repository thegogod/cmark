package markdown

import (
	"github.com/thegogod/cmark/html"
	"github.com/thegogod/cmark/tokens"
)

func (self *Markdown) ParseItalicAlt(parser html.Parser, ptr *tokens.Pointer) (html.Node, error) {
	return self.parseItalicAlt(parser, NewScanner(ptr))
}

func (self *Markdown) parseItalicAlt(parser html.Parser, scan *_Scanner) (*html.ItalicElement, error) {
	italic := html.I()

	if !scan.MatchCount(Underscore, 1) {
		return italic, scan.Curr().Error("expected '_'")
	}

	for !scan.Match(Underscore) {
		node, err := parser.ParseInline(scan.ptr)

		if node == nil {
			return italic, scan.Curr().Error("expected closing '_'")
		}

		if err != nil {
			return italic, err
		}

		italic.Push(node)
	}

	log.Debugln("italic_alt")
	return italic, nil
}
