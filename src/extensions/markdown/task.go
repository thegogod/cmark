package markdown

import (
	"github.com/google/uuid"

	"github.com/thegogod/cmark/html"
)

func (self *Markdown) parseTask(parser html.Parser, scan *_Scanner) (*html.LabelElement, error) {
	id := uuid.NewString()
	label := html.Label().WithFor(id)
	input := html.CheckBoxInput().WithId(id)

	if _, err := scan.Consume(LeftBracket, "expected '['"); err != nil {
		return label, err
	}

	checked, err := scan.Consume(Space, "expected ' ' or 'x'")

	if err != nil {
		checked, err = scan.Consume(Text, "expected ' ' or 'x'")

		if err != nil {
			return label, err
		}
	}

	if checked.String() != " " && checked.String() != "x" {
		return label, scan.Curr().Error("expected ' ' or 'x'")
	}

	if checked.String() == "x" {
		input.WithChecked(true)
	}

	if _, err = scan.Consume(RightBracket, "expected ']'"); err != nil {
		return label, err
	}

	if _, err = scan.Consume(Space, "expected ' '"); err != nil {
		return label, err
	}

	text := ""

	for !scan.Match(NewLine) {
		node, err := self.parseText(parser, scan)

		if err != nil {
			return label, err
		}

		if node == nil {
			break
		}

		text += string(node)
	}

	log.Debugln("task")
	label.Push(input)
	label.Push(html.Span(text))
	return label, nil
}
