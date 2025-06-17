package assert

type Test interface {
	Error(args ...any)
	Errorf(format string, args ...any)
	Fail()
	Fatal(args ...any)
	FailNow()
}
