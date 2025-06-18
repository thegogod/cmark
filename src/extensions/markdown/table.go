package markdown

import (
	"bytes"

	"github.com/thegogod/cmark/html"
	"github.com/thegogod/cmark/tokens"
)

func (self *Markdown) ParseTable(parser html.Parser, ptr *tokens.Pointer) (html.Node, error) {
	return self.parseTable(parser, NewScanner(ptr))
}

func (self *Markdown) parseTable(parser html.Parser, scan *Scanner) (*html.TableElement, error) {
	table := html.Table()

	if !scan.MatchCount(Pipe, 1) {
		return table, scan.Curr().Error("expected '|'")
	}

	columns := []*html.TableCellElement{}

	for !scan.Match(NewLine) {
		th := html.Th()

		for {
			if !scan.Match(Space) && !scan.Match(Tab) {
				break
			}
		}

		buff := []byte{}

		for !scan.Match(Pipe) {
			node, err := parser.ParseInline(scan.ptr)

			if node == nil {
				return table, scan.Curr().Error("invalid table header")
			}

			if err != nil {
				return table, err
			}

			if _, ok := node.(*html.BreakLineElement); ok {
				continue
			}

			if raw, ok := node.(html.Raw); ok && (string(raw) == " " || string(raw) == "\t") {
				buff = append(buff, raw...)
				continue
			}

			if len(buff) > 0 {
				th.Push(buff)
				buff = []byte{}
			}

			th.Push(node)
		}

		columns = append(columns, th)
	}

	if _, err := scan.Consume(Pipe, "expected opening '|'"); err != nil {
		return table, err
	}

	for i := range len(columns) {
		node, err := self.parseTextUntil(Pipe, parser, scan)

		if node == nil {
			return table, scan.Curr().Error("expected closing '|'")
		}

		if err != nil {
			return table, err
		}

		node = bytes.TrimSpace(node)
		dashes := 0

		for _, b := range node {
			if b == '-' {
				dashes++
			}

			if b != ':' && b != '-' {
				return table, scan.Curr().Error("invalid header separator")
			}
		}

		if dashes < 3 {
			return table, scan.Curr().Error("invalid header seperator")
		}

		leftAlign := bytes.HasPrefix(node, []byte{':'})
		rightAlign := bytes.HasSuffix(node, []byte{':'})

		if leftAlign && rightAlign {
			columns[i].SetStyle("text-align", "center")
		} else if leftAlign {
			columns[i].SetStyle("text-align", "left")
		} else if rightAlign {
			columns[i].SetStyle("text-align", "right")
		}
	}

	if _, err := scan.Consume(NewLine, "expected new line"); err != nil {
		return table, err
	}

	table.Push(html.THead(html.Tr(columns...)))
	rows := []*html.TableCellElement{}

	for scan.Match(Pipe) {
		for i := range len(columns) {
			td := html.Td()

			for {
				if !scan.Match(Space) && !scan.Match(Tab) {
					break
				}
			}

			buff := []byte{}

			for !scan.Match(Pipe) {
				node, err := parser.ParseInline(scan.ptr)

				if node == nil {
					return table, scan.Curr().Error("expected closing '|'")
				}

				if err != nil {
					return table, err
				}

				if _, ok := node.(*html.BreakLineElement); ok {
					continue
				}

				if raw, ok := node.(html.Raw); ok && (string(raw) == " " || string(raw) == "\t") {
					buff = append(buff, raw...)
					continue
				}

				if len(buff) > 0 {
					td.Push(buff)
					buff = []byte{}
				}

				td.Push(node)
			}

			rows = append(rows, td.WithStyles(columns[i].GetStyles()...))
		}

		if scan.Curr().Kind() == Eof {
			break
		}

		if _, err := scan.Consume(NewLine, "expected new line"); err != nil {
			return table, err
		}
	}

	log.Debugln("table")
	table.Push(html.TBody(html.Tr(rows...)))
	return table, nil
}
