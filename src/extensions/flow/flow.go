package flow

import (
	"fmt"

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
	log.Infoln(fmt.Sprintf("block => scope depth %d", self.scope.depth))

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
	log.Infoln(fmt.Sprintf("inline => scope depth %d", self.scope.depth))

	if scan.Match(Eof) {
		return nil, nil
	}

	if scan.Curr().Ext() != self.Name() {
		ptr.Back()
		scan.Next()
	}

	if _, err := scan.Consume(DoubleLeftBrace, "expected '{{'"); err != nil {
		return nil, err
	}

	statement, err := self.parseExpressionStatement(parser, scan)

	if err != nil {
		log.Warnln(err)
		return nil, err
	}

	if _, err := scan.Consume(DoubleRightBrace, "expected '}}'"); err != nil {
		return nil, err
	}

	if err := statement.Validate(self.scope); err != nil {
		return nil, err
	}

	return Html(statement), nil
}
