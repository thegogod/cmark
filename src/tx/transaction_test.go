package tx_test

import (
	"testing"

	"github.com/thegogod/cmark/tx"
)

func TestTransaction(t *testing.T) {
	t.Run("should rollback", func(t *testing.T) {
		a := "hello"
		tx := tx.New(&a)
		a = "world!"

		if a != "world!" {
			t.FailNow()
		}

		tx.Rollback()

		if a != "hello" {
			t.FailNow()
		}
	})

	t.Run("should rollback with pointers", func(t *testing.T) {
		type Test struct {
			A int
			B *string
		}

		b := "test"
		test := Test{B: &b}
		tx := tx.New(&test)

		test.A++
		*test.B = "123"
		tx.Rollback()

		if test.A != 0 || *test.B != "123" {
			t.FailNow()
		}
	})

	t.Run("should rollback compound", func(t *testing.T) {
		type A struct {
			A int
		}

		type B struct {
			B string
		}

		a := A{70}
		b := B{"hello world"}

		tx := tx.Compound(tx.New(&a), tx.New(&b))

		a.A = 90
		b.B = "lol"
		tx.Rollback()

		if a.A != 70 || b.B != "hello world" {
			t.FailNow()
		}
	})
}
