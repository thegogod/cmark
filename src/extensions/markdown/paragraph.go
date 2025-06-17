package markdown

import (
	"github.com/thegogod/cmark/html"
	"github.com/thegogod/cmark/tokens"
)

func (self *Markdown) ParseParagraph(parser html.Parser, ptr *tokens.Pointer) (html.Node, error) {
	return self.parseParagraph(parser, NewScanner(ptr))
}

func (self *Markdown) parseParagraph(parser html.Parser, scan *_Scanner) (*html.ParagraphElement, error) {
	paragraph := html.P()
	buff := html.Raw{}

	for scan.Curr().Kind() != Eof {
		node, err := parser.ParseInline(scan.ptr)

		if node == nil {
			break
		}

		if err != nil {
			return paragraph, err
		}

		if raw, ok := node.(html.Raw); ok && string(raw) == "\n" {
			if scan.Curr().Kind() == GreaterThan {
				break
			}

			buff = append(buff, raw...)

			if len(buff) > 1 {
				break
			}

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

	log.Debugln("paragraph")
	return paragraph, nil
}
