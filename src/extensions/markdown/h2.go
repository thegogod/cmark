package markdown

import (
	"github.com/thegogod/cmark/html"
	"github.com/thegogod/cmark/tokens"
)

func (self *Markdown) ParseH2(parser html.Parser, ptr *tokens.Pointer) (html.Node, error) {
	return self.parseH2(parser, NewScanner(ptr))
}

func (self *Markdown) parseH2(parser html.Parser, scan *_Scanner) (*html.HeadingElement, error) {
	if !scan.MatchCount(Hash, 2) || !scan.Match(Space) {
		return nil, scan.Curr().Error("expected '## '")
	}

	log.Debugln("h2")
	heading := html.H2()

	for scan.Curr().Kind() != Eof && scan.Curr().Kind() != NewLine {
		node, err := parser.ParseInline(scan.ptr)

		if node == nil || err != nil {
			return heading, err
		}

		heading.Push(node)
	}

	return heading, nil
}
