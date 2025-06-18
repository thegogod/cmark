package frontmatter

import (
	"github.com/goccy/go-yaml"
	"github.com/thegogod/cmark/extensions/markdown"
	"github.com/thegogod/cmark/html"
	"github.com/thegogod/cmark/tokens"
)

func (self *Frontmatter) ParseMetaData(parser html.Parser, ptr *tokens.Pointer) (html.Node, error) {
	return self.parseMetaData(parser, markdown.NewScanner(ptr))
}

func (self *Frontmatter) parseMetaData(parser html.Parser, scan *markdown.Scanner) (html.MetaDataElement, error) {
	el := html.MetaData()

	if !(scan.Ptr().Sof() && scan.MatchCount(markdown.Dash, 3)) {
		return el, scan.Curr().Error("expected '---'")
	}

	if _, err := scan.Consume(markdown.NewLine, "expected newline"); err != nil {
		return el, err
	}

	data := []byte{}

	for {
		line, err := self.parseTextUntil(markdown.NewLine, parser, scan)

		if line == nil {
			return el, scan.Curr().Error("expected newline at end of key value pair")
		}

		if err != nil {
			return el, err
		}

		data = append(data, line...)

		if scan.MatchCount(markdown.Dash, 3) {
			break
		}

		data = append(data, '\n')
	}

	if _, err := scan.Consume(markdown.NewLine, "expected newline"); err != nil {
		return el, err
	}

	if _, err := scan.Consume(markdown.NewLine, "expected newline"); err != nil {
		return el, err
	}

	if err := yaml.Unmarshal(data, &el); err != nil {
		return el, err
	}

	return el, nil
}
