package markdown

import (
	"github.com/thegogod/cmark/html"
	"github.com/thegogod/cmark/tokens"
)

func (self *Markdown) ParseH5(parser html.Parser, ptr *tokens.Pointer) (html.Node, error) {
	return self.parseH5(parser, NewScanner(ptr))
}

func (self *Markdown) parseH5(parser html.Parser, scan *_Scanner) (*html.HeadingElement, error) {
	heading := html.H5()

	if !(scan.MatchCount(Hash, 5) && scan.Match(Space)) {
		return heading, scan.Curr().Error("expected '##### '")
	}

	for scan.Curr().Kind() != Eof && scan.Curr().Kind() != NewLine {
		node, err := parser.ParseInline(scan.ptr)

		if node == nil || err != nil {
			return heading, err
		}

		heading.Push(node)
	}

	log.Debugln("h5")
	return heading, nil
}
