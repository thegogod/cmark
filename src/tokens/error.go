package tokens

import (
	"fmt"
	"strconv"
)

type Error struct {
	Start   Position
	End     Position
	Message string
}

func Err(start Position, end Position, message string) *Error {
	return &Error{
		Start:   start,
		End:     end,
		Message: message,
	}
}

func (self Error) Error() string {
	return self.String()
}

func (self Error) String() string {
	line := strconv.Itoa(self.Start.Ln)

	if self.End.Ln != self.Start.Ln {
		line = fmt.Sprintf("%d-%d", self.Start.Ln, self.End.Ln)
	}

	column := strconv.Itoa(self.Start.Col)

	if self.End.Col != self.Start.Col {
		column = fmt.Sprintf("%d-%d", self.Start.Col, self.End.Col)
	}

	return fmt.Sprintf(
		"[%s:%s] => %s",
		line,
		column,
		self.Message,
	)
}
