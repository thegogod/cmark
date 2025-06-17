package markdown

import (
	"github.com/thegogod/cmark/html"
	"github.com/thegogod/cmark/tokens"
)

func (self *Markdown) ParseImage(parser html.Parser, ptr *tokens.Pointer) (html.Node, error) {
	return self.parseImage(parser, NewScanner(ptr))
}

func (self *Markdown) parseImage(parser html.Parser, scan *_Scanner) (*html.ImageElement, error) {
	image := html.Img()

	if !scan.Match(Bang) {
		return image, scan.Curr().Error("expected '!'")
	}

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

	log.Debugln("image")
	image.WithSrc(string(node))
	return image, nil
}
