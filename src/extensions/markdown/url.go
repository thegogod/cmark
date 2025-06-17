package markdown

import (
	"fmt"
	"strings"

	"github.com/thegogod/cmark/html"
	"github.com/thegogod/cmark/tokens"
	"github.com/thegogod/cmark/tx"
)

func (self *Markdown) ParseUrl(parser html.Parser, ptr *tokens.Pointer) (html.Node, error) {
	return self.parseUrl(parser, NewScanner(ptr))
}

func (self *Markdown) parseUrl(parser html.Parser, scan *_Scanner) (*html.AnchorElement, error) {
	tx := tx.New(scan)
	text, err := self.parseText(parser, scan)

	if text == nil || err != nil {
		defer tx.Rollback()
		return nil, scan.Curr().Error("expected text")
	}

	if !strings.HasPrefix(string(text), "http") {
		defer tx.Rollback()
		return nil, scan.Curr().Error("expected 'http' prefix")
	}

	link := html.A()
	protocol, _ := self.parseText(parser, scan)

	if _, err := scan.Consume(Colon, "expected ':'"); err != nil {
		return link, err
	}

	if _, err := scan.Consume(Slash, "expected '/'"); err != nil {
		return link, err
	}

	if _, err := scan.Consume(Slash, "expected '/'"); err != nil {
		return link, err
	}

	path := []byte{}

	for {
		part, err := self.parseText(parser, scan)

		if part == nil || err != nil {
			return link, err
		}

		path = append(path, part...)

		if scan.Curr().Kind() != Period {
			break
		}
	}

	log.Debugln("url")
	url := fmt.Sprintf("%s://%s", protocol, path)
	link.WithHref(url)
	link.Push(url)
	return link, nil
}
