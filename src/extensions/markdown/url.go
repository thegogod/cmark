package markdown

import (
	"fmt"
	"strings"

	"github.com/thegogod/cmark/html"
	"github.com/thegogod/cmark/tokens"
)

func (self *Markdown) ParseUrl(parser html.Parser, ptr *tokens.Pointer) (html.Node, error) {
	return self.parseUrl(parser, NewScanner(ptr))
}

func (self *Markdown) parseUrl(parser html.Parser, scan *Scanner) (*html.AnchorElement, error) {
	link := html.A()
	protocol, err := self.parseText(parser, scan)

	if protocol == nil || err != nil {
		return link, scan.Curr().Error("expected text")
	}

	if !strings.HasPrefix(string(protocol), "http") {
		return link, scan.Curr().Error("expected 'http' prefix")
	}

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
