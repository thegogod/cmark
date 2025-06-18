package markdown

import (
	"bytes"

	"github.com/thegogod/cmark/html"
	"github.com/thegogod/cmark/logging"
	"github.com/thegogod/cmark/tokens"
	"github.com/thegogod/cmark/tx"
)

var log = logging.Console("cmark.markdown")

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

func (self *Markdown) ParseBlock(parser html.Parser, ptr *tokens.Pointer) (html.Node, error) {
	return self.parseBlock(parser, NewScanner(ptr))
}

func (self *Markdown) ParseInline(parser html.Parser, ptr *tokens.Pointer) (html.Node, error) {
	return self.parseInline(parser, NewScanner(ptr))
}

func (self *Markdown) ParseText(parser html.Parser, ptr *tokens.Pointer) ([]byte, error) {
	return self.parseText(parser, NewScanner(ptr))
}

func (self *Markdown) ParseTextUntil(kind rune, parser html.Parser, ptr *tokens.Pointer) ([]byte, error) {
	return self.parseTextUntil(kind, parser, NewScanner(ptr))
}

func (self *Markdown) parseBlock(parser html.Parser, scan *Scanner) (html.Node, error) {
	if scan.Match(Eof) {
		return nil, nil
	}

	var node html.Node = nil
	var err error = nil

	for range self.blockQuoteDepth - 1 {
		if !scan.Match(GreaterThan) {
			break
		}
	}

	scan.NextWhile(NewLine)
	tx := tx.Compound(tx.New(self), tx.New(scan.ptr))
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
		node, err = self.parseCodeBlock(parser, scan)
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

	log.Debugln("block")
	return node, err
}

func (self *Markdown) parseInline(parser html.Parser, scan *Scanner) (html.Node, error) {
	if scan.Match(Eof) {
		return nil, nil
	}

	var node html.Node = nil
	var err error = nil

	tx := tx.Compound(tx.New(self), tx.New(scan.ptr))
	node, err = self.parseHtml(parser, scan)

	if err != nil {
		tx.Rollback()
		node, err = self.parseBold(parser, scan)
	}

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

	if err != nil {
		tx.Rollback()
	}

	if node == nil || err != nil {
		log.Debugln("text")
		text, texterr := self.parseText(parser, scan)

		if text != nil {
			node = html.Raw(text)
		}

		err = texterr
	}

	log.Debugln("inline", scan.ptr.Start.String(), scan.ptr.End.String())
	return node, err
}

func (self *Markdown) parseText(_ html.Parser, scan *Scanner) ([]byte, error) {
	if scan.Curr().Kind() == Eof {
		return nil, nil
	}

	text := html.Raw(scan.Curr().Bytes())
	scan.Next()

	if bytes.Equal(text, []byte{' '}) {
		return text, nil
	}

	for scan.Curr().Kind() == Text {
		text = append(text, scan.Curr().Bytes()...)
		scan.Next()
	}

	return text, nil
}

func (self *Markdown) parseTextUntil(kind rune, parser html.Parser, scan *Scanner) ([]byte, error) {
	if scan.Curr().Kind() == Eof {
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
