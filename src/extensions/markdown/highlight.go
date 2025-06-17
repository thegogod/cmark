package markdown

import (
	"github.com/thegogod/cmark/html"
	"github.com/thegogod/cmark/tokens"
)

func (self *Markdown) ParseHighlight(parser html.Parser, ptr *tokens.Pointer) (html.Node, error) {
	return self.parseHighlight(parser, NewScanner(ptr))
}

func (self *Markdown) parseHighlight(parser html.Parser, scan *_Scanner) (*html.MarkElement, error) {
	mark := html.Mark()

	if !scan.MatchCount(EqualsEquals, 1) {
		return mark, scan.Curr().Error("expected '=='")
	}

	for !scan.MatchCount(EqualsEquals, 1) {
		node, err := parser.ParseInline(scan.ptr)

		if node == nil {
			return mark, scan.Curr().Error("expected closing '=='")
		}

		if err != nil {
			return mark, err
		}

		mark.Push(node)
	}

	log.Debugln("highlight")
	return mark, nil
}
