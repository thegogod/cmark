package assert

type Assertion interface {
	Assert(t Test)
	AssertNow(t Test)
}

type Evaluation interface {
	Evaluate(value any) error
}
