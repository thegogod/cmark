package flow_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/thegogod/cmark/extensions/flow"
	"github.com/thegogod/cmark/tokens"
)

func TestScanner(t *testing.T) {
	t.SkipNow()

	t.Run("should scan", func(t *testing.T) {
		data, err := os.ReadFile(filepath.Join("..", "markdown", "testcases", "text.md"))

		if err != nil {
			t.Fatal(err)
		}

		ptr := tokens.Ptr(data)
		scanner := flow.NewScanner(ptr)

		for {
			token, err := scanner.Scan()

			if token == nil || token.Kind() == flow.Eof {
				break
			}

			if err != nil {
				t.Fatal(err)
			}

			t.Logf("%v => %s", token.Kind(), token.String())
		}
	})
}
