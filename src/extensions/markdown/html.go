package markdown

import (
	"bytes"

	"github.com/thegogod/cmark/html"
	"github.com/thegogod/cmark/tokens"
	"github.com/thegogod/cmark/tx"
)

func (self *Markdown) ParseHtml(parser html.Parser, ptr *tokens.Pointer) (html.Node, error) {
	return self.parseHtml(parser, NewScanner(ptr))
}

func (self *Markdown) parseHtml(parser html.Parser, scan *_Scanner) (*html.Element, error) {
	if self.path == nil {
		self.path = []string{}
	}

	log.Debugln("html")
	name := []byte{}
	scan.NextWhile(Space, Tab)

	for scan.Match(Text) || scan.Match(Underscore) || scan.Match(Dash) {
		name = append(name, scan.Prev().Bytes()...)
	}

	el := html.Elem(string(name))
	self.path = append(self.path, string(name))
	depth := len(self.path)

	for scan.NextWhile(Space, Tab) > 0 && scan.Curr().Kind() == Text {
		attr := []byte{}
		value := []byte{}

		for scan.Match(Text) || scan.Match(Underscore) || scan.Match(Dash) {
			attr = append(attr, scan.Prev().Bytes()...)
		}

		if scan.Match(Equals) {
			if _, err := scan.Consume(DoubleQuote, "expected '\"'"); err != nil {
				return el, err
			}

			v, err := self.parseTextUntil(DoubleQuote, parser, scan)

			if v == nil {
				return el, scan.Curr().Error("expected closing '\"'")
			}

			if err != nil {
				return el, err
			}

			value = v
		}

		el.SetAttr(string(attr), string(value))
	}

	isVoid := false

	if scan.Curr().Kind() == Slash {
		isVoid = true
		el.Void()
		scan.Next()
	}

	if _, err := scan.Consume(GreaterThan, "expected closing '>'"); err != nil {
		return el, err
	}

	if !isVoid {
		for {
			scan.NextWhile(NewLine, Tab)

			if self.parseClosingTag(scan, name, depth) {
				break
			}

			content, err := parser.ParseInline(scan.ptr)

			if content == nil {
				return el, scan.Curr().Error("expected closing tag")
			}

			if err != nil {
				return el, err
			}

			el.Push(content)
		}
	}

	self.path = self.path[:len(self.path)-1]
	return el, nil
}

func (self *Markdown) parseClosingTag(scan *_Scanner, name []byte, depth int) bool {
	if !scan.Match(LessThan, Slash) {
		return false
	}

	tx := tx.New(scan)
	scan.NextWhile(Space, Tab)
	tag := []byte{}

	for scan.Match(Text) || scan.Match(Underscore) || scan.Match(Dash) {
		tag = append(tag, scan.Prev().Bytes()...)
	}

	if !bytes.Equal(tag, name) {
		tx.Rollback()
		return false
	}

	scan.NextWhile(Space, Tab)

	if !scan.Match(GreaterThan) || depth != len(self.path) {
		tx.Rollback()
		return false
	}

	return true
}
