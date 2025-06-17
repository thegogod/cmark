package markdown

import (
	"github.com/thegogod/cmark/emojis"
	"github.com/thegogod/cmark/html"
	"github.com/thegogod/cmark/tokens"
)

func (self *Markdown) ParseEmoji(parser html.Parser, ptr *tokens.Pointer) (html.Node, error) {
	return self.parseEmoji(parser, NewScanner(ptr))
}

func (self *Markdown) parseEmoji(parser html.Parser, scan *_Scanner) (html.Raw, error) {
	alias := html.Raw{}

	if !scan.MatchCount(Colon, 1) {
		return alias, scan.Curr().Error("expected ':'")
	}

	for !scan.Match(Colon) {
		node, err := self.parseText(parser, scan)

		if node == nil {
			return alias, scan.Curr().Error("expected closing ':'")
		}

		if err != nil {
			return alias, err
		}

		alias = append(alias, node...)
	}

	emoji, exists := emojis.Get(string(alias))

	if !exists {
		return alias, scan.Curr().Error("emoji alias not found")
	}

	log.Debugln("emoji")
	return html.Raw(emoji.Emoji), nil
}
