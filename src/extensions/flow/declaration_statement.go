package flow

import (
	"github.com/thegogod/cmark/html"
	"github.com/thegogod/cmark/tokens"
)

func (self *Flow) ParseDeclarationStatement(parser html.Parser, ptr *tokens.Pointer) (html.Node, error) {
	node, err := self.parseDeclarationStatement(parser, NewScanner(ptr))

	if err != nil {
		return nil, err
	}

	return Html(node), nil
}

func (self *Flow) parseDeclarationStatement(parser html.Parser, scan *Scanner) (Statement, error) {
	log.Infoln("current token =>", scan.Curr().Kind(), scan.Curr().String())

	if scan.Match(Let) || scan.Match(Const) {
		return self.parseVariableStatement(parser, scan)
	}

	return self.parseStatement(parser, scan)
}
