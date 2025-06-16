package markdown

import (
	"github.com/thegogod/cmark/ast"
	"github.com/thegogod/cmark/emojis"
	"github.com/thegogod/cmark/html"
	"github.com/thegogod/cmark/tokens"
)

func (self *Markdown) ParseEmoji(parser ast.Parser, ptr *tokens.Pointer) (ast.Node, error) {
	return self.parseEmoji(parser, NewScanner(ptr))
}

func (self *Markdown) parseEmoji(parser ast.Parser, scan *_Scanner) (html.Raw, error) {
	if !scan.MatchCount(Colon, 1) {
		return nil, scan.curr.Error("expected ':'")
	}

	alias := html.Raw{}

	for !scan.Match(Colon) {
		node, err := self.parseText(parser, scan)

		if node == nil {
			return alias, scan.curr.Error("expected closing ':'")
		}

		if err != nil {
			return alias, err
		}

		alias = append(alias, node...)
	}

	emoji, exists := emojis.Get(string(alias))

	if !exists {
		return alias, scan.curr.Error("emoji alias not found")
	}

	return html.Raw(emoji.Emoji), nil
}
