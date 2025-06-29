package markdown

import (
	"fmt"

	"github.com/thegogod/cmark/html"
	"github.com/thegogod/cmark/tokens"
)

func (self *Markdown) ParseCodeBlock(parser html.Parser, ptr *tokens.Pointer) (html.Node, error) {
	return self.parseCodeBlock(parser, NewScanner(ptr))
}

func (self *Markdown) parseCodeBlock(parser html.Parser, scan *Scanner) (*html.PreElement, error) {
	code := html.Code()

	if !scan.MatchCount(BackQuote, 3) {
		return html.Pre(code), scan.Curr().Error("expected '```'")
	}

	lang, err := self.parseTextUntil(NewLine, parser, scan)

	if lang == nil || err != nil {
		return html.Pre(code), err
	}

	if len(lang) > 0 {
		code.AddClass(fmt.Sprintf("language-%s", lang))
	}

	buff := html.Raw{}

	for !scan.MatchCount(BackQuote, 3) {
		node, err := self.parseText(parser, scan)

		if node == nil {
			return html.Pre(code), scan.Curr().Error("expected closing '```'")
		}

		if err != nil {
			return html.Pre(code), err
		}

		if string(node) == "\n" {
			buff = append(buff, node...)
			continue
		}

		if len(buff) > 0 {
			code.Push(buff)
			buff = html.Raw{}
		}

		code.Push(node)
	}

	log.Debugln("code_block")
	return html.Pre(code), nil
}
