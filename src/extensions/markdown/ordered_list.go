package markdown

import (
	"github.com/thegogod/cmark/html"
	"github.com/thegogod/cmark/tokens"
)

func (self *Markdown) ParseOrderedList(parser html.Parser, ptr *tokens.Pointer) (html.Node, error) {
	return self.parseOrderedList(parser, NewScanner(ptr))
}

func (self *Markdown) parseOrderedList(parser html.Parser, scan *Scanner) (*html.OrderedListElement, error) {
	ol := html.Ol()

	if !(scan.Match(Integer) && scan.Match(Period) && scan.Match(Space)) {
		return nil, scan.Curr().Error("expected '{int}. '")
	}

	self.listDepth++

	for {
		node, err := self.parseListItem(parser, scan)

		if node == nil || err != nil {
			self.listDepth--
			return ol, err
		}

		ol.Push(node)

		if !scan.MatchCount(Tab, self.listDepth-1) {
			break
		}

		if !(scan.Match(Integer) && scan.Match(Period) && scan.Match(Space)) {
			break
		}
	}

	log.Debugln("ordered_list")
	self.listDepth--
	return ol, nil
}
