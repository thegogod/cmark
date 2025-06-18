package markdown

import (
	"github.com/thegogod/cmark/html"
	"github.com/thegogod/cmark/tokens"
)

func (self *Markdown) ParseUnorderedList(parser html.Parser, ptr *tokens.Pointer) (html.Node, error) {
	return self.parseUnorderedList(parser, NewScanner(ptr))
}

func (self *Markdown) parseUnorderedList(parser html.Parser, scan *Scanner) (*html.UnorderedListElement, error) {
	ul := html.Ul()

	if !(scan.Match(Dash) && scan.Match(Space)) {
		return ul, scan.Curr().Error("expected '- '")
	}

	self.listDepth++

	for {
		node, err := self.parseListItem(parser, scan)

		if node == nil || err != nil {
			self.listDepth--
			return ul, err
		}

		ul.Push(node)

		if !scan.MatchCount(Tab, self.listDepth-1) {
			break
		}

		if !(scan.Match(Dash) && scan.Match(Space)) {
			break
		}
	}

	log.Debugln("unordered_list")
	self.listDepth--
	return ul, nil
}
