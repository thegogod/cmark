package markdown

import (
	"github.com/thegogod/cmark/ast"
	"github.com/thegogod/cmark/html"
	"github.com/thegogod/cmark/tokens"
)

func (self *Markdown) ParseImage(parser ast.Parser, ptr *tokens.Pointer) (ast.Node, error) {
	return self.parseImage(parser, NewScanner(ptr))
}

func (self *Markdown) parseImage(parser ast.Parser, scan *_Scanner) (*html.ImageElement, error) {
	if !scan.Match(Bang) {
		return nil, scan.Curr().Error("expected '!'")
	}

	image := html.Img()

	if _, err := scan.Consume(LeftBracket, "expected '['"); err != nil {
		return image, err
	}

	node, err := self.parseTextUntil(RightBracket, parser, scan)

	if node == nil || err != nil {
		return image, err
	}

	image.WithAlt(string(node))

	if _, err := scan.Consume(LeftParen, "expected '('"); err != nil {
		return image, err
	}

	node, err = self.parseTextUntil(RightParen, parser, scan)

	if node == nil || err != nil {
		return image, err
	}

	image.WithSrc(string(node))
	return image, nil
}
