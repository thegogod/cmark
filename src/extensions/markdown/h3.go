package markdown

import (
	"github.com/thegogod/cmark/html"
	"github.com/thegogod/cmark/tokens"
)

func (self *Markdown) ParseH3(parser html.Parser, ptr *tokens.Pointer) (html.Node, error) {
	return self.parseH3(parser, NewScanner(ptr))
}

func (self *Markdown) parseH3(parser html.Parser, scan *Scanner) (*html.HeadingElement, error) {
	heading := html.H3()

	if !(scan.MatchCount(Hash, 3) && scan.Match(Space)) {
		return heading, scan.Curr().Error("expected '### '")
	}

	for scan.Curr().Kind() != Eof && scan.Curr().Kind() != NewLine {
		node, err := parser.ParseInline(scan.ptr)

		if node == nil || err != nil {
			return heading, err
		}

		heading.Push(node)
	}

	log.Debugln("h3")
	return heading, nil
}
