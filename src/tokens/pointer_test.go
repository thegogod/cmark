package tokens_test

import (
	"testing"

	"github.com/thegogod/cmark/assert"
	"github.com/thegogod/cmark/tokens"
	"github.com/thegogod/cmark/tx"
)

func TestPointer(t *testing.T) {
	t.Run("should rollback", func(t *testing.T) {
		ptr := tokens.Ptr([]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g'})
		ptr.Next()
		tx := tx.New(ptr)

		assert.Expect(string(ptr.Curr())).Equal("a").AssertNow(t)
		assert.Expect(string(ptr.Peek())).Equal("b").AssertNow(t)

		for !ptr.Eof() && ptr.Peek() != 'd' {
			ptr.Next()
		}

		token := ptr.Ok(2)
		assert.Expect(string(ptr.Peek())).Equal("d").AssertNow(t)
		assert.Expect(token.String()).Equal("abc").AssertNow(t)
		assert.Expect(ptr.Iter.Curr).Equal(token).AssertNow(t)

		tx.Rollback()
		assert.Expect(string(ptr.Curr())).Equal("a").AssertNow(t)
		assert.Expect(string(ptr.Peek())).Equal("b").AssertNow(t)
		assert.Expect(int(ptr.Iter.Curr.Kind())).Equal(-1).AssertNow(t)
	})
}
