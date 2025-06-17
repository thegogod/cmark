package assert_test

import (
	"testing"

	"github.com/thegogod/cmark/assert"
)

type Tester struct {
	failed bool
	now    bool
	args   []any
}

func (self *Tester) Error(args ...any) {
	self.failed = true
	self.args = args
}

func (self *Tester) Errorf(format string, args ...any) {
	self.failed = true
	self.args = args
}

func (self *Tester) Fail() {
	self.failed = true
}

func (self *Tester) Fatal(args ...any) {
	self.failed = true
	self.now = true
	self.args = args
}

func (self *Tester) FailNow() {
	self.failed = true
	self.now = true
}

func TestExpect(t *testing.T) {
	t.Run("should assert", func(t *testing.T) {
		tester := &Tester{}
		expect := assert.Expect(1).Equal(2)
		expect.Assert(tester)

		if !tester.failed || tester.now {
			t.FailNow()
		}
	})

	t.Run("should assert now", func(t *testing.T) {
		tester := &Tester{}
		expect := assert.Expect("testing").Equal("testing123")
		expect.AssertNow(tester)

		if !tester.failed || !tester.now {
			t.FailNow()
		}
	})

	t.Run("should assert with message", func(t *testing.T) {
		tester := &Tester{}
		expect := assert.Expect(1).Equal(2).Message("a test message")
		expect.Assert(tester)

		if !tester.failed || tester.args == nil || len(tester.args) != 1 {
			t.FailNow()
		}

		if tester.args[0] != "a test message" {
			t.FailNow()
		}
	})

	t.Run("and or", func(t *testing.T) {
		tester := &Tester{}
		expect := assert.Expect("test").And(assert.Equal(1), assert.Equal("test"))
		expect.Assert(tester)

		if !tester.failed {
			t.FailNow()
		}

		tester = &Tester{}
		expect = assert.Expect("test").Or(assert.Equal(1), assert.Equal("test"))
		expect.Assert(tester)

		if tester.failed {
			t.FailNow()
		}
	})

	t.Run("len", func(t *testing.T) {
		t.Run("should succeed", func(t *testing.T) {
			tester := &Tester{}
			expect := assert.Expect("test").Len(4)
			expect.Assert(tester)

			if tester.failed {
				t.FailNow()
			}
		})

		t.Run("should fail with invalid length", func(t *testing.T) {
			tester := &Tester{}
			expect := assert.Expect("test").Len(3)
			expect.Assert(tester)

			if !tester.failed {
				t.FailNow()
			}
		})

		t.Run("should fail with invalid type", func(t *testing.T) {
			tester := &Tester{}
			expect := assert.Expect(1).Len(3)
			expect.Assert(tester)

			if !tester.failed {
				t.FailNow()
			}
		})
	})

	t.Run("nil", func(t *testing.T) {
		t.Run("should succeed", func(t *testing.T) {
			tester := &Tester{}
			expect := assert.Expect(nil).Nil()
			expect.Assert(tester)

			if tester.failed {
				t.FailNow()
			}
		})

		t.Run("should fail", func(t *testing.T) {
			tester := &Tester{}
			expect := assert.Expect("test").Nil()
			expect.Assert(tester)

			if !tester.failed {
				t.FailNow()
			}
		})
	})
}
