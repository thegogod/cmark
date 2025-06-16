package markdown

import (
	"github.com/thegogod/cmark/ast"
	"github.com/thegogod/cmark/html"
	"github.com/thegogod/cmark/tokens"
)

func (self *Markdown) ParseParagraph(parser ast.Parser, ptr *tokens.Pointer) (ast.Node, error) {
	return self.parseParagraph(parser, NewScanner(ptr))
}

func (self *Markdown) parseParagraph(parser ast.Parser, scan *_Scanner) (*html.ParagraphElement, error) {
	paragraph := html.P()
	buff := html.Raw{}

	for scan.curr.Kind() != Eof {
		node, err := parser.ParseInline(scan.ptr)

		if node == nil {
			break
		}

		if err != nil {
			return paragraph, err
		}

		if raw, ok := node.(html.Raw); ok && string(raw) == "\n" {
			if scan.curr.Kind() == GreaterThan {
				break
			}

			buff = append(buff, raw...)
			continue
		}

		if len(buff) > 0 {
			paragraph.Push(buff)
			buff = html.Raw{}
		}

		paragraph.Push(node)
	}

	if len(paragraph.Children()) == 0 {
		return nil, nil
	}

	return paragraph, nil
}
