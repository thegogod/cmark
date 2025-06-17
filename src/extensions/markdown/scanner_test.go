package markdown_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/thegogod/cmark/extensions/markdown"
	"github.com/thegogod/cmark/tokens"
)

func TestScanner(t *testing.T) {
	t.SkipNow()

	t.Run("should scan", func(t *testing.T) {
		data, err := os.ReadFile(filepath.Join("testcases", "block_quote.md"))

		if err != nil {
			t.Fatal(err)
		}

		ptr := tokens.Ptr(data)
		scanner := markdown.NewScanner(ptr)

		for {
			token, err := scanner.Scan()

			if token == nil || token.Kind() == markdown.Eof {
				break
			}

			if err != nil {
				t.Fatal(err)
			}

			t.Logf("%v => %s", token.Kind(), token.String())
		}
	})
}
