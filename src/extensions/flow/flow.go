package flow

import (
	"github.com/thegogod/cmark/html"
	"github.com/thegogod/cmark/logging"
	"github.com/thegogod/cmark/tokens"
)

var log = logging.Console("cmark.flow")

type Flow struct {
	scope *Scope
}

func New() *Flow {
	return &Flow{
		scope: NewScope(),
	}
}

func (self Flow) Name() string {
	return "flow"
}

func (self *Flow) ParseBlock(parser html.Parser, ptr *tokens.Pointer) (html.Node, error) {
	scan := NewScanner(ptr)

	if scan.Match(Eof) {
		return nil, nil
	}

	if scan.Curr().Ext() != self.Name() {
		ptr.Back()
		scan.Next()
	}

	statement, err := self.parseDeclarationStatement(parser, scan)

	if err != nil {
		log.Warnln(err)
		return nil, err
	}

	if err := statement.Validate(self.scope); err != nil {
		return nil, err
	}

	return Html(statement), nil
}

func (self *Flow) ParseInline(parser html.Parser, ptr *tokens.Pointer) (html.Node, error) {
	scan := NewScanner(ptr)

	if scan.Match(Eof) {
		return nil, nil
	}

	if scan.Curr().Ext() != self.Name() {
		ptr.Back()
		scan.Next()
	}

	statement, err := self.parseExpressionStatement(parser, scan)

	if err != nil {
		log.Warnln(err)
		return nil, err
	}

	if err := statement.Validate(self.scope); err != nil {
		return nil, err
	}

	return Html(statement), nil
}
