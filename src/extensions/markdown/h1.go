package markdown

import (
	"github.com/thegogod/cmark/html"
	"github.com/thegogod/cmark/tokens"
)

func (self *Markdown) ParseH1(parser html.Parser, ptr *tokens.Pointer) (html.Node, error) {
	return self.parseH1(parser, NewScanner(ptr))
}

func (self *Markdown) parseH1(parser html.Parser, scan *_Scanner) (*html.HeadingElement, error) {
	if !scan.MatchCount(Hash, 1) || !scan.Match(Space) {
		return nil, scan.Curr().Error("expected '# '")
	}

	heading := html.H1()

	for scan.Curr().Kind() != Eof && scan.Curr().Kind() != NewLine {
		node, err := parser.ParseInline(scan.ptr)

		if node == nil || err != nil {
			return heading, err
		}

		heading.Push(node)
	}

	return heading, nil
}
