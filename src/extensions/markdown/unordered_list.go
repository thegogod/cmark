package markdown

import (
	"github.com/thegogod/cmark/html"
	"github.com/thegogod/cmark/tokens"
)

func (self *Markdown) ParseUnorderedList(parser html.Parser, ptr *tokens.Pointer) (html.Node, error) {
	return self.parseUnorderedList(parser, NewScanner(ptr))
}

func (self *Markdown) parseUnorderedList(parser html.Parser, scan *_Scanner) (*html.UnorderedListElement, error) {
	if !scan.Match(Dash) || !scan.Match(Space) {
		return nil, scan.Curr().Error("expected '- '")
	}

	log.Debugln("unordered_list")
	self.listDepth++
	ul := html.Ul()

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

	self.listDepth--
	return ul, nil
}
