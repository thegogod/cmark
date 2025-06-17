package markdown

import (
	"github.com/thegogod/cmark/html"
	"github.com/thegogod/cmark/tokens"
	"github.com/thegogod/cmark/tx"
)

func (self *Markdown) ParseListItem(parser html.Parser, ptr *tokens.Pointer) (html.Node, error) {
	return self.parseListItem(parser, NewScanner(ptr))
}

func (self *Markdown) parseListItem(parser html.Parser, scan *_Scanner) (*html.ListItemElement, error) {
	li := html.Li()
	t := tx.New(scan)
	node, err := self.parseTask(parser, scan)

	if err == nil && node != nil {
		li.Push(node)
		return li, nil
	}

	t.Rollback()

	for scan.Curr().Kind() != Eof {
		node, err := parser.ParseInline(scan.ptr)

		if err != nil {
			t.Rollback()
			return li, err
		}

		if node == nil {
			break
		}

		if raw, ok := node.(html.Raw); ok && string(raw) == "\n" {
			if !scan.MatchCount(Tab, self.listDepth) {
				break
			}

			node, err = nil, nil
			tx := tx.New(scan)

			if scan.Match(Integer) && scan.Match(Period) && scan.Match(Space) {
				node, err = self.parseOrderedList(parser, scan)
			} else if scan.Match(Dash) && scan.Match(Space) {
				node, err = self.parseUnorderedList(parser, scan)
			}

			if node != nil && err == nil {
				li.Push(node)
			} else {
				tx.Rollback()
			}

			break
		}

		li.Push(node)
	}

	return li, nil
}
