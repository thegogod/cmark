package markdown

import (
	"bytes"

	"github.com/thegogod/cmark/ast"
	"github.com/thegogod/cmark/html"
	"github.com/thegogod/cmark/tokens"
	"github.com/thegogod/cmark/tx"
)

type Markdown struct {
	blockQuoteDepth int
	listDepth       int
	path            []string
}

func New() *Markdown {
	return &Markdown{path: []string{}}
}

func (self *Markdown) Name() string {
	return "markdown"
}

func (self *Markdown) ParseBlock(parser ast.Parser, ptr *tokens.Pointer) (ast.Node, error) {
	return self.parseBlock(parser, NewScanner(ptr))
}

func (self *Markdown) ParseInline(parser ast.Parser, ptr *tokens.Pointer) (ast.Node, error) {
	return self.parseInline(parser, NewScanner(ptr))
}

func (self *Markdown) ParseSyntax(parser ast.Parser, ptr *tokens.Pointer, name string) (ast.Node, error) {
	return nil, nil
}

func (self *Markdown) ParseText(parser ast.Parser, ptr *tokens.Pointer) ([]byte, error) {
	return self.parseText(parser, NewScanner(ptr))
}

func (self *Markdown) ParseTextUntil(kind rune, parser ast.Parser, ptr *tokens.Pointer) ([]byte, error) {
	return self.parseTextUntil(kind, parser, NewScanner(ptr))
}

func (self *Markdown) parseBlock(parser ast.Parser, scan *_Scanner) (ast.Node, error) {
	if scan.Match(Eof) {
		return nil, nil
	}

	var node ast.Node = nil
	var err error = nil

	for range self.blockQuoteDepth - 1 {
		if !scan.Match(GreaterThan) {
			break
		}
	}

	if scan.Match(NewLine) {
		return self.parseBlock(parser, scan)
	}

	tx := tx.New(scan)
	node, err = self.parseHtml(parser, scan)

	if err != nil {
		tx.Rollback()
		node, err = self.parseH1(parser, scan)
	}

	if err != nil {
		tx.Rollback()
		node, err = self.parseH2(parser, scan)
	}

	if err != nil {
		tx.Rollback()
		node, err = self.parseH3(parser, scan)
	}

	if err != nil {
		tx.Rollback()
		node, err = self.parseH4(parser, scan)
	}

	if err != nil {
		tx.Rollback()
		node, err = self.parseH5(parser, scan)
	}

	if err != nil {
		tx.Rollback()
		node, err = self.parseH6(parser, scan)
	}

	if err != nil {
		tx.Rollback()
		node, err = self.parseHorizontalRule(parser, scan)
	}

	if err != nil {
		tx.Rollback()
		node, err = self.parseBlockQuote(parser, scan)
	}

	if err != nil {
		tx.Rollback()
		node, err = self.parseUnorderedList(parser, scan)
	}

	if err != nil {
		tx.Rollback()
		node, err = self.parseOrderedList(parser, scan)
	}

	if err != nil {
		tx.Rollback()
		node, err = self.parseTable(parser, scan)
	}

	if err != nil {
		tx.Rollback()
		node, err = self.parseParagraph(parser, scan)
	}

	return node, err
}

func (self *Markdown) parseInline(parser ast.Parser, scan *_Scanner) (ast.Node, error) {
	if scan.Match(Eof) {
		return nil, nil
	}

	var node ast.Node = nil
	var err error = nil

	tx := tx.New(scan)
	node, err = self.parseBold(parser, scan)

	if err != nil {
		tx.Rollback()
		node, err = self.parseBoldAlt(parser, scan)
	}

	if err != nil {
		tx.Rollback()
		node, err = self.parseItalic(parser, scan)
	}

	if err != nil {
		tx.Rollback()
		node, err = self.parseItalicAlt(parser, scan)
	}

	if err != nil {
		tx.Rollback()
		node, err = self.parseStrike(parser, scan)
	}

	if err != nil {
		tx.Rollback()
		node, err = self.parseStrikeAlt(parser, scan)
	}

	if err != nil {
		tx.Rollback()
		node, err = self.parseHighlight(parser, scan)
	}

	if err != nil {
		tx.Rollback()
		node, err = self.parseCode(parser, scan)
	}

	if err != nil {
		tx.Rollback()
		node, err = self.parseEmoji(parser, scan)
	}

	if err != nil {
		tx.Rollback()
		node, err = self.parseLink(parser, scan)
	}

	if err != nil {
		tx.Rollback()
		node, err = self.parseUrl(parser, scan)
	}

	if err != nil {
		tx.Rollback()
		node, err = self.parseImage(parser, scan)
	}

	if err != nil {
		tx.Rollback()
		node, err = self.parseBreakLine(parser, scan)
	}

	if err != nil {
		tx.Rollback()
		node, err = self.parseNewLine(parser, scan)
	}

	if node == nil || err != nil {
		text, texterr := self.parseText(parser, scan)

		if text != nil {
			node = html.Raw(text)
		}

		err = texterr
	}

	return node, err
}

func (self *Markdown) parseText(_ ast.Parser, scan *_Scanner) ([]byte, error) {
	if scan.curr.Kind() == Eof {
		return nil, nil
	}

	text := html.Raw(scan.curr.Bytes())
	scan.Next()

	if bytes.Equal(text, []byte{' '}) {
		return text, nil
	}

	for scan.curr.Kind() == Text {
		text = append(text, scan.curr.Bytes()...)
		scan.Next()
	}

	return text, nil
}

func (self *Markdown) parseTextUntil(kind rune, parser ast.Parser, scan *_Scanner) ([]byte, error) {
	if scan.curr.Kind() == Eof {
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
