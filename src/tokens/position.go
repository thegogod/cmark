package tokens

import "fmt"

type Position struct {
	Ln    int
	Col   int
	Index int
}

func (self Position) String() string {
	return fmt.Sprintf(
		"%d:%d",
		self.Ln+1,
		self.Col+1,
	)
}
