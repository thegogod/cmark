package frontmatter

import (
	"bytes"

	"github.com/thegogod/cmark/extensions/markdown"
	"github.com/thegogod/cmark/html"
	"github.com/thegogod/cmark/logging"
	"github.com/thegogod/cmark/tokens"
)

var log = logging.Console("cmark.frontmatter")

type Frontmatter struct {
}

func New() *Frontmatter {
	return &Frontmatter{}
}

func (self Frontmatter) Name() string {
	return "frontmatter"
}

func (self *Frontmatter) ParseBlock(parser html.Parser, ptr *tokens.Pointer) (html.Node, error) {
	return self.parseBlock(parser, markdown.NewScanner(ptr))
}

func (self *Frontmatter) ParseInline(parser html.Parser, ptr *tokens.Pointer) (html.Node, error) {
	return nil, nil
}

func (self *Frontmatter) ParseText(parser html.Parser, ptr *tokens.Pointer) ([]byte, error) {
	return self.parseText(parser, markdown.NewScanner(ptr))
}

func (self *Frontmatter) ParseTextUntil(kind rune, parser html.Parser, ptr *tokens.Pointer) ([]byte, error) {
	return self.parseTextUntil(kind, parser, markdown.NewScanner(ptr))
}

func (self *Frontmatter) parseBlock(parser html.Parser, scan *markdown.Scanner) (html.Node, error) {
	if scan.Match(markdown.Eof) {
		return nil, nil
	}

	log.Debugln("block")
	return self.parseMetaData(parser, scan)
}

func (self *Frontmatter) parseText(_ html.Parser, scan *markdown.Scanner) ([]byte, error) {
	if scan.Curr().Kind() == markdown.Eof {
		return nil, nil
	}

	text := html.Raw(scan.Curr().Bytes())
	scan.Next()

	if bytes.Equal(text, []byte{' '}) {
		return text, nil
	}

	for scan.Curr().Kind() == markdown.Text {
		text = append(text, scan.Curr().Bytes()...)
		scan.Next()
	}

	return text, nil
}

func (self *Frontmatter) parseTextUntil(kind rune, parser html.Parser, scan *markdown.Scanner) ([]byte, error) {
	if scan.Curr().Kind() == markdown.Eof {
		return nil, nil
	}

	text := html.Raw{}

	for !scan.Match(kind) {
		node, err := self.parseText(parser, scan)

		if node == nil || err != nil {
			return text, err
		}

		text = append(text, node...)
	}

	return text, nil
}
